package test

import (
	"testing"

	"github.com/cayohollanda/go_splunk"
)

var conn = &go_splunk.SplunkConnection{
	APIURL:   "https://localhost:8089",
	Username: "admin",
	Password: "password",
}

func TestSearch(t *testing.T) {
	resp, err := conn.GetSearchResults("index=test")

	if err != nil {
		t.Error("Returns error on request")
	}

	// solve this problem in test (not converting JSON to Struct)
	if resp.Result.Raw == "" {
		t.Error("Returns []string slice of raws equals 0")
	}
}
