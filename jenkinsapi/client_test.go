package jenkinsapi

import (
	"strings"
	"testing"
)

func Test_NewClient_BadURL(t *testing.T) {
	client, err := NewClient(":")
	if err == nil || err.Error() != "parse :: missing protocol scheme" {
		t.Error("NewClient did not fail when passed an invalid URL")
	}

	if (client != nil) {
		t.Error("NewClient returned a client despite having an error")
	}
}

func Test_NewClient_NoScheme(t *testing.T) {
	client, err := NewClient("/meh")
	if err == nil || !strings.HasPrefix(err.Error(), "Could not determine scheme from url:") {
		t.Error("NewClient did not fail when passed a URL with no host")
	}

	if (client != nil) {
		t.Error("NewClient returned a client despite having an error")
	}
}

func Test_NewClient_NoHost(t *testing.T) {
	client, err := NewClient("http://")
	if err == nil || !strings.HasPrefix(err.Error(), "Could not determine host from url:") {
		t.Error("NewClient did not fail when passed a URL with no host")
	}

	if (client != nil) {
		t.Error("NewClient returned a client despite having an error")
	}
}

func Test_NewClient_Success(t *testing.T) {
	client, err := NewClient("http://localhost")
	if (err != nil) {
		t.Error("NewClient returned an error despite having a valid url: " + err.Error())
	}

	if (client == nil) {
		t.Error("NewClient returned no client despite having no errors")
	}
}
