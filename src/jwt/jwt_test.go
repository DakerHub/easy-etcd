package jwt

import (
	"testing"
)

const username = "test"
const pwd = "pwd"

func TestEncodeJWT(t *testing.T) {
	token, err := EncodeJWT(username, pwd)

	if err != nil {
		t.Fatal(err)
	}
	if token == "" {
		t.Fatal("No token return")
	}
}

func TestDecodeJWT(t *testing.T) {
	token, _ := EncodeJWT(username, pwd)
	username2, pwd2, err := DecodeJWT(token)

	if err != nil {
		t.Fatal(err)
	}

	if username2 != username && pwd2 != pwd {
		t.Fatal("Fail to match")
	}
}
