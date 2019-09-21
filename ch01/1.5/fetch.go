// fetch prints the conten found in a URL
// from this section, execise and example code are in one folder
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		url = CheckURL(url)
		rsp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Http get from[%s] failed:%s\n", url, err)
			os.Exit(1)
		}

		// example in gopl book
		// readAll(rsp)

		// execise1.7
		Copy(rsp)
	}
}

func readAll(rsp *http.Response) {

	b, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close() // bad design?

	if err != nil {
		fmt.Fprintf(os.Stderr, "Read http response body failed:%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", b)
	PrintHTTPStatus(rsp)
}
