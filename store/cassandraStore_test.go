package store

import "testing"

func TestCassandra(t *testing.T) {
	Init()
}

func TestSave(t *testing.T) {
	Save("longlongUrl", "xxx1", "127.0.0.1")
}
