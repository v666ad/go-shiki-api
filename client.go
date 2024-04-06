package shikimori

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/v666ad/go-shiki-api/types"
)

var (
	ShikiSchema = "https"
	ShikiDomain = "shikimori.one"
	UserAgent   = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/115.0"
)

const (
	TypeTopic = "Topic"
	TypeUser  = "User"
)

type Client struct {
	*types.Me
	Cookies    string
	XCsrfToken string
	client     http.Client
}

func NewClient(cookies string, xCsrfToken string) (*Client, error) {
	var client http.Client
	var err error

	client.Timeout = 30 * time.Second

	shikiClient := &Client{
		Cookies:    cookies,
		XCsrfToken: xCsrfToken,
		client:     client,
	}

	shikiClient.Me, err = shikiClient.GetMe()
	if err != nil {
		return nil, err
	}

	return shikiClient, nil
}

func (c *Client) MakeRequest(method string, path string, urlParams url.Values, data io.Reader) (*http.Response, error) {
	var (
		req *http.Request
		err error
	)
	if urlParams != nil {
		req, err = http.NewRequest(method, ShikiSchema+"://"+ShikiDomain+"/"+path+"?"+urlParams.Encode(), data)
	} else {
		req, err = http.NewRequest(method, ShikiSchema+"://"+ShikiDomain+"/"+path, data)
	}
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Cookie", c.Cookies)
	req.Header.Set("X-CSRF-Token", c.XCsrfToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		switch resp.StatusCode {
		case 404:
			return nil, ErrNotFound
		case 429:
			return nil, ErrTooManyRequests
		}
		return nil, errors.New("bad status " + resp.Request.Method + " " + req.URL.String() + " -> " + resp.Status)
	}

	return resp, err
}
