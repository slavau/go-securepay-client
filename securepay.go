package securepay

import (
	"net/http"
	"net/url"
)
// A Client manages communication with the SecurePay XML API.
type Client struct {
	client *http.Client
	baseURL *url.URL

	// Services used for talking to different parts of the JIRA API.
	Payment          *PaymentService
}

func NewClient(httpClient *http.Client, baseURL string) (*Client, error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		client: httpClient,
		baseURL: parsedBaseURL,
	}

	c.Payment = &PaymentService{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	// TODO
	return nil, nil
}

func (c *Client) Perform(req *http.Request, v interface{}) (*Response, error) {
	// TODO
	return nil, nil
}

