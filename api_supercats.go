package cex

// SuperCat or "super category" represents a general purpose category, there are
// only a few of these. The following supercats exist at the time of writing:
//
// - Gaming (1)
// - Film & TV (2)
// - Computing (3)
// - Phones (4)
// - Electronics (5)
// - Music (8)
//
type SuperCat struct {
	SuperCatID           int    `json:"superCatId"`
	SuperCatFriendlyName string `json:"superCatFriendlyName"`
}

// SuperCats simple returns a list of all super categories
func (client *Client) SuperCats() (result []SuperCat, err error) {
	payload, err := client.request("supercats", nil)
	if err != nil {
		return
	}
	return payload.Data.SuperCats, nil
}
