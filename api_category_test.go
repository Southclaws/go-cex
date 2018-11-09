package cex

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestClient_Categories(t *testing.T) {
	tests := []struct {
		productLines []int
	}{
		{[]int{91}},
		{[]int{54}},
		{[]int{17}},
		{[]int{78}},
		{[]int{51, 23, 68}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt), func(t *testing.T) {
			gotResult, err := client.Categories(tt.productLines)
			if err != nil {
				t.Error(err)
			}
			pretty.Println(gotResult) //nolint:errcheck
		})
	}
}
