package main

import (
	"io"
	"net/http"
	"os"
)

// Copy do execise 1.7
func Copy(rsp *http.Response) {
	io.Copy(os.Stdout, rsp.Body)
	PrintHTTPStatus(rsp)
}
