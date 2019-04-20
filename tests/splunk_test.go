package test

import (
	"testing"

	"github.com/cayohollanda/go_splunk"
)

var conn = &go_splunk.SplunkConnection{
	APIURL:   "https://localhost:8089",
	Username: "admin",
	Password: "admin@1234",
}

func TestSearch(t *testing.T) {
	_, err := conn.GetSearchResults("index=test")

	if err != nil {
		t.Error("Returns error on request")
	}
}
