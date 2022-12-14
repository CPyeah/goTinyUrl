package store

import (
	"errors"
	"github.com/gocql/gocql"
	"goTinyUrl/cache"
	"log"
	"time"
)

var cassandra *gocql.Session

func init() {
	Init()
}

func Init() {
	cluster := gocql.NewCluster("127.0.0.1:9042", "127.0.0.1:9043")
	//cluster.Keyspace = "tinyUrl_service"
	cluster.Consistency = gocql.Quorum
	cluster.ProtoVersion = 4
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.DCAwareRoundRobinPolicy("localDC"))
	session, err := cluster.CreateSession()
	cassandra = session
	simpleHandleErr(err)

	// create keyspace
	err = session.Query("CREATE KEYSPACE IF NOT EXISTS tinyUrl_service WITH REPLICATION = {'class': 'SimpleStrategy', 'replication_factor': 2};").Exec()
	simpleHandleErr(err)

	// use keyspace
	err = session.Query("use tinyUrl_service;").Exec()

	// create table
	createTableCql := "create table IF NOT EXISTS tinyUrl_service.tiny_url (short_url varchar primary key,  origin_url  varchar, ip  varchar,   create_time timestamp);"
	err = session.Query(createTableCql).Exec()
	simpleHandleErr(err)
}

func simpleHandleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Save(longUrl string, shortUrl string, ip string) (error, bool) {
	// save to Cassandra
	insertCql := "insert into tinyUrl_service.tiny_url (short_url, origin_url, ip, create_time) VALUES (?, ?, ?, ?) IF NOT EXISTS;"
	var shortUrlCAS string
	var longUrlCAS string
	var ipCAS string
	var createTimeCAS int64
	ok, err := cassandra.Query(insertCql,
		shortUrl, longUrl, ip, time.Now().UnixNano()).ScanCAS(&shortUrlCAS, &longUrlCAS, &ipCAS, &createTimeCAS)
	if err != nil {
		log.Println(err)
		return err, ok
	}
	//add to bloomFilter
	cache.Add(shortUrl)

	return nil, ok
}

func Get(shortUrl string) (string, error) {
	// check from bloomFilter
	if !cache.Exists(shortUrl) {
		return "", errors.New(shortUrl + " not exists")
	}

	var originUrl string
	// find from cache
	originUrl = cache.Get(shortUrl)
	if len(originUrl) > 0 {
		return originUrl, nil
	}
	// get from Cassandra
	selectCql := "select origin_url from tinyUrl_service.tiny_url where short_url = ?"
	err := cassandra.Query(selectCql, shortUrl).Scan(&originUrl)
	if err != nil {
		return "", err
	}

	// cache result
	cache.Sava(shortUrl, originUrl)

	return originUrl, nil
}
