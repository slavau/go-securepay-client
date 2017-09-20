package securepay

import (
	"bytes"
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
)

// A Client manages communication with the SecurePay XML API.
type Client struct {
	client  *http.Client
	baseURL *url.URL

	// Services used for talking to different parts of the JIRA API.
	Payment *PaymentService
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
		client:  httpClient,
		baseURL: parsedBaseURL,
	}

	c.Payment = &PaymentService{client: c}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}
	newURL := c.baseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := xml.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, newURL.String(), buf)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/xml")
	return req, nil
}

func (c *Client) Perform(req *http.Request) (*http.Response, error) {
	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
