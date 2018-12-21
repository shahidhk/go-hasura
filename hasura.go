package hasura

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Client can execute queries against an endpoint
type Client struct {
	Endpoint string
	Headers  map[string]string
	client   *http.Client
}

// NewClient returns a Client for given endpoint and headers
func NewClient(endpoint string, headers map[string]string) *Client {
	return &Client{
		Endpoint: endpoint,
		Headers:  headers,
		client:   &http.Client{},
	}
}

// Execute executes the Query r using the Client c and returns an error
// Response data can be unmarshalled to the passed interface
func (c *Client) Execute(r Query, response interface{}) error {
	payload, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// check if status code is non 200
	if res.StatusCode != 200 {
		// it's an error response, decode the error
		var e Error
		err = json.NewDecoder(res.Body).Decode(&e)
		if err != nil {
			return err
		}
		return e
	}

	// response is 200, decode response into the passed interface
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	return nil
}
