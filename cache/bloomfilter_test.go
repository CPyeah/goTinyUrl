package cache

import (
	"fmt"
	"testing"
)

func TestNewRedisBloomFilter(t *testing.T) {
	bloom := NewRedisBloomFilter(Rdb, "test", 10000)

	//bs := []byte("hello")
	//_ = bloom.Put(bs)
	exists, err := bloom.MightContain([]byte("hello"))
	//exists, err = bloom.MightContain([]byte("hello1"))
	fmt.Println(exists, err)
}
