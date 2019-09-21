package main

import (
	"strings"
)

// CheckURL do 1.8
func CheckURL(url string) string {
	if strings.HasPrefix(url, "http://") == false {
		url = string("http://") + url
	}

	return url
}
