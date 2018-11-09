package cex

import (
	"os"
	"testing"
)

var client *Client

func TestMain(m *testing.M) {
	client = NewClient()
	client.http.SetDebug(true)

	ret := m.Run()
	os.Exit(ret)
}
