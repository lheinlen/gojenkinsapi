package jenkinsapi

import (
	"strings"
	"testing"
)

func Test_FromURL_BadURL(t *testing.T) {
	client, err := FromURL(":")
	if err == nil || err.Error() != "parse :: missing protocol scheme" {
		t.Error("FromURL did not fail when passed an invalid URL")
	}

	if client != nil {
		t.Error("FromURL returned a client despite having an error")
	}
}

func Test_FromURL_NoScheme(t *testing.T) {
	client, err := FromURL("/meh")
	if err == nil || !strings.HasPrefix(err.Error(), "Scheme (http/https) is required") {
		t.Error("FromURL did not fail when passed a URL with no scheme")
	}

	if client != nil {
		t.Error("FromURL returned a client despite having an error")
	}
}

func Test_FromURL_UnknownScheme(t *testing.T) {
	client, err := FromURL("zttp://example.com")
	if err == nil || !strings.HasPrefix(err.Error(), "Unknown scheme: zttp") {
		t.Error("FromURL did not fail when passed a URL with an unknown scheme")
	}

	if client != nil {
		t.Error("FromURL returned a client despite having an error")
	}
}

func Test_FromURL_NoHost(t *testing.T) {
	client, err := FromURL("http://")
	if err == nil || !strings.HasPrefix(err.Error(), "Host is required") {
		t.Error("FromURL did not fail when passed a URL with no host")
	}

	if client != nil {
		t.Error("FromURL returned a client despite having an error")
	}
}

func Test_FromURL_HttpsSuccess(t *testing.T) {
	client, err := FromURL("https://example.com")
	if err != nil {
		t.Error("FromURL returned an error despite having a valid url: " + err.Error())
	}

	if client.Scheme != "https" {
		t.Error("Scheme was not properly parsed from the URL")
	}
}

func Test_FromURL_SimpleSuccess(t *testing.T) {
	client, err := FromURL("http://example.com")
	if err != nil {
		t.Error("FromURL returned an error despite having a valid url: " + err.Error())
	}

	if client == nil {
		t.Error("FromURL returned no client despite having no errors")
	}

	if client.Scheme != "http" {
		t.Error("Scheme was not properly parsed from the URL")
	}

	if client.Host != "example.com" {
		t.Error("Host was not properly parsed from the URL")
	}

	if client.Username != "" {
		t.Error("User was not properly parsed from the URL")
	}

	if client.Password != "" {
		t.Error("Password was not properly parsed from the URL")
	}
}

func Test_FromURL_FullSuccess(t *testing.T) {
	client, err := FromURL("http://user:pass@example.com:80")
	if err != nil {
		t.Error("FromURL returned an error despite having a valid url: " + err.Error())
	}

	if client == nil {
		t.Error("FromURL returned no client despite having no errors")
	}

	if client.Scheme != "http" {
		t.Error("Scheme was not properly parsed from the URL")
	}

	if client.Host != "example.com:80" {
		t.Error("Host was not properly parsed from the URL")
	}

	if client.Username != "user" {
		t.Error("User was not properly parsed from the URL")
	}

	if client.Password != "pass" {
		t.Error("Password was not properly parsed from the URL")
	}
}
