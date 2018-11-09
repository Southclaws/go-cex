package cex

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestClient_ProductLines(t *testing.T) {
	tests := []struct {
		superCatIds []int
	}{
		{[]int{1}},
		{[]int{2}},
		{[]int{3}},
		{[]int{4}},
		{[]int{5}},
		{[]int{8}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			gotResult, err := client.ProductLines(tt.superCatIds)
			if err != nil {
				t.Error(err)
			}
			pretty.Println(gotResult) //nolint:errcheck
		})
	}
}
