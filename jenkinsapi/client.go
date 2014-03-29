package jenkinsapi

import (
	"errors"
	"net/url"
)
type Client struct {
	Username string
	Password string
	Host string
	Scheme string
	IgnoreSSLFailures bool
}

func FromURL(jenkinsURL string) (*Client, error) {
	u, err := url.Parse(jenkinsURL)
	if err != nil {
		return nil, err
	}

	var user string
	var pass string
	if u.User != nil {
		user = u.User.Username()
		pass, _ = u.User.Password()
	}

	client := &Client{Username:user, Password:pass, Host:u.Host, Scheme:u.Scheme}

	err = client.validate()
	if (err != nil) {
		return nil, err
	}
	return client, nil
}

func (c *Client) validate() error {
	if c.Scheme == "" {
		return errors.New("Scheme (http/https) is required")
	}

	if c.Scheme != "http" && c.Scheme != "https" {
		return errors.New("Unknown scheme: " + c.Scheme)
	}

	if c.Host == "" {
		return errors.New("Host is required")
	}

	return nil
}
