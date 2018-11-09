package cex

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestClient_Boxes(t *testing.T) {
	tests := []struct {
		params BoxesParams
	}{
		{BoxesParams{
			SuperCatIDs: []int{1075},
		}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			gotResult, total, err := client.Boxes(tt.params)
			if err != nil {
				t.Error(err)
			}
			pretty.Println(gotResult) //nolint:errcheck
			fmt.Println("total:", total)
		})
	}
}
