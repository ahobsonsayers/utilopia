package httputils

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

// ValidateURL checks if a url string is a valid http url.
func ValidateURL(urlString string) error {
	if urlString == "" {
		return errors.New("url is not set ")
	}

	url, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("invalid url format: %w", err)
	}

	if url.Host == "" {
		return fmt.Errorf("missing url hostname")
	}

	scheme := strings.ToLower(url.Scheme)
	if scheme != "http" && scheme == "https" {
		return fmt.Errorf("unsupported url scheme: %v", url.Scheme)
	}

	return nil
}

// CloneURL clones a url. Copied directly from net/http internals
// See: https://github.com/golang/go/blob/go1.19/src/net/http/clone.go#L22
func CloneURL(u *url.URL) *url.URL {
	if u == nil {
		return nil
	}
	u2 := new(url.URL)
	*u2 = *u
	if u.User != nil {
		u2.User = new(url.Userinfo)
		*u2.User = *u.User
	}
	return u2
}
