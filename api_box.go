package cex

import (
	"encoding/json"

	"github.com/Southclaws/qstring"
)

// Box represents an individual product
type Box struct {
	BoxID                string    `json:"boxId"`
	BoxName              string    `json:"boxName"`
	IsMasterBox          int       `json:"isMasterBox"`
	CategoryID           int       `json:"categoryId"`
	CategoryName         string    `json:"categoryName"`
	CategoryFriendlyName string    `json:"categoryFriendlyName"`
	SuperCatID           int       `json:"superCatId"`
	SuperCatName         string    `json:"superCatName"`
	SuperCatFriendlyName string    `json:"superCatFriendlyName"`
	ImageURLs            ImageURLs `json:"imageUrls"`
	CannotBuy            int       `json:"cannotBuy"`
	IsNewBox             int       `json:"isNewBox"`
	SellPrice            float64   `json:"sellPrice"`
	CashPrice            float64   `json:"cashPrice"`
	ExchangePrice        float64   `json:"exchangePrice"`
	BoxRating            float64   `json:"boxRating"`
	OutOfStock           int       `json:"outOfStock"`
	OutOfEcomStock       int       `json:"outOfEcomStock"`
}

// ImageURLs contains a simple set of image URLs
type ImageURLs struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
}

// SortBy defines the different sort types
type SortBy string

// The different sort types
var (
	SortByMostPopular SortBy = "relevance" //
	SortBySellPrice   SortBy = "sellprice" //
	SortByName        SortBy = "boxname"   //
	SortByRating      SortBy = "rating"    //
)

// BoxesParams represents the URL query parameters for the Boxes endpoint
type BoxesParams struct {
	// One of these must be present
	SuperCatIDs []int `qstring:"superCatIds"` // super category ID
	CategoryIDs []int `qstring:"categoryIds"` // category ID

	// These are all optional
	FirstRecord int    `qstring:"firstRecord,omitempty"` // default: 50, basically a database OFFSET
	Count       int    `qstring:"count,omitempty"`       // default: 1, and a database LIMIT
	SortBy      string `qstring:"sortBy,omitempty"`      // default: relevance, attribute to sort by
	SortOrder   string `qstring:"sortOrder,omitempty"`   // default: desc, sort order
}

// Boxes returns product boxes from the given search parameters
func (client *Client) Boxes(params BoxesParams) (result []Box, err error) {
	q, err := qstring.Marshal(&params)
	if err != nil {
		return
	}

	// override the ID fields because they use weird JSON array syntax
	if len(params.SuperCatIDs) > 0 {
		var superCatIDList []byte
		superCatIDList, err = json.Marshal(params.SuperCatIDs)
		if err != nil {
			return nil, err
		}
		q.Set("superCatIds", string(superCatIDList))
	}
	if len(params.CategoryIDs) > 0 {
		var categoryIDList []byte
		categoryIDList, err = json.Marshal(params.CategoryIDs)
		if err != nil {
			return nil, err
		}
		q.Set("categoryIds", string(categoryIDList))
	}

	payload, err := client.request("boxes", q)
	if err != nil {
		return
	}

	return payload.Data.Boxes, nil
}
