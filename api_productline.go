package cex

import (
	"encoding/json"

	"github.com/pkg/errors"
)

// ProductLine represents a top-level line of product categories
type ProductLine struct {
	SuperCatID      int    `json:"superCatId"`
	ProductLineID   int    `json:"productLineId"`
	ProductLineName string `json:"productLineName"`
	TotalCategories int    `json:"totalCategories"`
}

// ProductLines returns a list of product lines given a list of superCatIds
func (client *Client) ProductLines(superCatIds []int) (result []ProductLine, err error) {
	var superCatList []byte
	superCatList, err = json.Marshal(superCatIds)
	if err != nil {
		return
	}

	var wrapped ResponseWrapper
	resp, err := client.http.R().
		SetQueryParam("superCatIds", string(superCatList)).
		SetResult(&wrapped).
		Get("productlines")
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

	result = wrapped.Response.Data.ProductLines

	return
}
