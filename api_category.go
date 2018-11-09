package cex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// Category represents a category of boxes with a super category
type Category struct {
	SuperCatID           int    `json:"superCatId"`
	CategoryID           int    `json:"categoryId"`
	CategoryFriendlyName string `json:"categoryFriendlyName"`
	ProductLineID        int    `json:"productLineId"`
	TotalBoxes           int    `json:"totalBoxes"`
}

// Categories returns product categories from for a set of productlines
func (client *Client) Categories(productLines []int) (result []Category, err error) {
	var productLineList []byte
	productLineList, err = json.Marshal(productLines)
	if err != nil {
		return
	}

	var wrapped ResponseWrapper
	resp, err := client.http.R().
		SetQueryParam("productLineIds", string(productLineList)).
		SetResult(&wrapped).
		Get("categories")
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = errors.Errorf("endpoint responded with non-OK status: %s (%d)", resp.Status(), resp.StatusCode())
		return
	}

	if wrapped.Response.Ack != "Success" {
		err = errors.Errorf("response ack was not 'Success': '%s', error: %s (%s)",
			wrapped.Response.Ack,
			wrapped.Response.Error.InternalMessage,
			wrapped.Response.Error.Code)
		return
	}

	result = wrapped.Response.Data.Categories
	return
}
