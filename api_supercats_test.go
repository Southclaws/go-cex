package cex

import (
	"testing"

	"github.com/kr/pretty"
)

func TestClient_SuperCats(t *testing.T) {
	supercats, err := client.SuperCats()
	if err != nil {
		t.Error(err)
	}
	pretty.Println(supercats) //nolint:errcheck
}
