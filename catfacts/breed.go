package catfacts

func (c *Client) ListCatBreeds() (Breeds, error) {
	req, err := c.newRequest("GET", "/breeds", nil)
	if err != nil {
		return Breeds{}, err
	}
	var b Breeds
	_, err = c.do(req, &b)
	return b, err
}

type Breeds struct {
	PagedResponse
	Data []struct {
		Breed   string `json:"breed"`
		Country string `json:"country"`
		Origin  string `json:"origin"`
		Coat    string `json:"coat"`
		Pattern string `json:"pattern"`
	} `json:"data"`
}
