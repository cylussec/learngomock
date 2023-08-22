package catfacts

import "fmt"

// Will get a single, random Cat Fact(r) from catfacts.ninja
// Returns a single CatFact type, as well as an error if the API returns an error
func (c *Client) ListCatFact() (CatFact, error) {
	req, err := c.newRequest("GET", "/fact", nil)
	if err != nil {
		return CatFact{}, err
	}
	var cf CatFact
	_, err = c.do(req, &cf)
	return cf, err
}

func (c *Client) ListCatFacts(page int) (CatFacts, error) {
	url := fmt.Sprintf("/facts?page=%d", page)
	req, err := c.newRequest("GET", url, nil)
	if err != nil {
		return CatFacts{}, err
	}
	var cf CatFacts
	_, err = c.do(req, &cf)
	return cf, err
}

// Gets a specific number of cat facts, specified by the argument num
// Returns a pointer to a slice of Cat Facts of the specified number
func (c *Client) GetNumberOfCatFacts(num int) (*[]CatFact, error) {
	page := 1
	var cf []CatFact
	for {
		ret, err := c.ListCatFacts(page)
		if err != nil {
			return nil, err
		}
		cf = append(cf, ret.Data...)
		if len(cf) > num {
			break
		}

	}
	return &cf, nil
}

// Cat Fact return struct
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func (cf *CatFact) String() string {
	return cf.Fact
}

// Full http response of a slice of CatFacts
type CatFacts struct {
	PagedResponse
	Data []CatFact `json:"data"`
}
