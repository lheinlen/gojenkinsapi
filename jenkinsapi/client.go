package jenkinsapi

import (
	"errors"
	"net/url"
)
type Client struct {
	BaseUrl *url.URL
}

func NewClient(baseUrl string) (*Client, error) {
	url, err := url.Parse(baseUrl)
	if (err != nil) {
		return nil, err
	}

	if (url.Scheme == "") {
		return nil, errors.New("Could not determine scheme from url: " + baseUrl)
	}

	if (url.Host == "") {
		return nil, errors.New("Could not determine host from url: " + baseUrl)
	}

	client := &Client{BaseUrl:url}
	return client, nil
}
