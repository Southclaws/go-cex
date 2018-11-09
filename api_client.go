package cex

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
	"gopkg.in/resty.v1"
)

// Client represents a HTTP client for interacting with CeX web resources.
type Client struct {
	http *resty.Client
}

// NewClient creates a new CeX RESTful client.
func NewClient() *Client {
	client := Client{}

	client.http = resty.New().
		SetHostURL("https://wss2.cex.uk.webuy.io/v3/").
		SetRESTMode()

	return &client
}

// ResponseWrapper wraps each response from the API.
type ResponseWrapper struct {
	Response ResponsePayload `json:"response"`
}

// ResponsePayload is the only member of the response object. Why it's wrapped
// in this way? Who knows... Maybe another version of the API contained more
// fields in the parent object.
type ResponsePayload struct {
	Ack   string        `json:"ack"`
	Data  ResponseData  `json:"data"`
	Error ResponseError `json:"error"`
}

// ResponseData contains the actual data from a query. Each query fills a
// different field.
type ResponseData struct {
	// These are the main query result fields
	SuperCats    []SuperCat    `json:"superCats"`
	ProductLines []ProductLine `json:"productLines"`
	Categories   []Category    `json:"categories"`
	Boxes        []Box         `json:"boxes"`

	// These are additional pieces of data sent back with query results
	TotalRecords int     `json:"totalRecords"`
	MinPrice     float64 `json:"minPrice"`
	MaxPrice     float64 `json:"maxPrice"`
	Facets       Facets  `json:"facets"`
}

// Facets seems to be some sort of summary of a set of query results containing
// data such as the number of items from each manufacturer etc. It hasn't been
// fully explored and modelled yet.
type Facets struct {
	SuperCatName           []map[string]interface{} // TODO
	CategoryFriendlyName   []map[string]interface{} // TODO
	ManufacturerName       []map[string]interface{} // TODO
	NetworkName            []map[string]interface{} // TODO
	AttributeStructureInfo interface{}              // ???
	AttributeInfo          interface{}              // ???
}

// ResponseError contains the structured error data (if any) in a response.
type ResponseError struct {
	Code            string        `json:"code"`
	InternalMessage string        `json:"internal_message"`
	MoreInfo        []interface{} `json:"moreInfo"`
}

func (re ResponseError) Error() string {
	return fmt.Sprintf("%s (%s) %v", re.InternalMessage, re.Code, re.MoreInfo)
}

// request is a generic request function used by all API callers
func (client *Client) request(endpoint string, query url.Values) (payload ResponsePayload, err error) {
	var wrapped ResponseWrapper
	resp, err := client.http.R().
		SetMultiValueQueryParams(query).
		SetResult(&wrapped).
		Get(endpoint)
	if err != nil {
		return
	}

	if wrapped.Response.Ack == "" {
		err = errors.Errorf("%s (%d)", resp.Status(), resp.StatusCode())
	} else if wrapped.Response.Ack != "Success" {
		err = wrapped.Response.Error
	} else {
		payload = wrapped.Response
	}

	return
}
