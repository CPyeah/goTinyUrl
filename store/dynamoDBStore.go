package store

func Save(longUrl string, shortUrl string, ip string) error {
	// save to dynamoDB

	// add to bloomFilter

	return nil
}

func Get(shortUrl string) string {
	// check from bloomFilter

	// find from cache

	// get from dynamoDB

	// cache result

	return "https://github.com/CPyeah/goTinyUrl"
}
