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
