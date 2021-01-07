package etcd

import (
	"testing"
)

var m = []string{"http://localhost:2379"}

var auth = Auth{
	Password: "",
	Username: "",
}

func TestConnent(t *testing.T) {
	m := []string{"http://localhost:2379"}
	_, err := Connect(m, auth)

	if err != nil {
		t.Fatal(err)
	}
}

func TestPut(t *testing.T) {
	_, err := Put(m, auth, "/test2/sub", "xxxxx")

	if err != nil {
		t.Fatal(err)
	}
}

func TestGet(t *testing.T) {
	res, err := GetSingle(m, auth, "/test")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.Count)
	t.Log(res.Kvs)
}

func TestGetByPrefix(t *testing.T) {
	res, err := GetByPrefix(m, auth, "")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(res.Count)
	t.Log(res.Kvs)
}

func TestDel(t *testing.T) {
	err := Del(m, auth, "\265\353-\377\034q")
	if err != nil {
		t.Fatal(err)
	}
}
