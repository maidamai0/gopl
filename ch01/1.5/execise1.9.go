package main

import (
	"fmt"
	"net/http"
)

// PrintHTTPStatus do 1.9
func PrintHTTPStatus(rsp *http.Response) {
	fmt.Printf("status:%s", rsp.Status)
}
