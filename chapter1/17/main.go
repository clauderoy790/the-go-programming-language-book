package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !hasValidPrefix(url) {
			url = fmt.Sprintf("http://%s", url)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nfetch: statusCode:%d,  reading %s: %v\n", resp.StatusCode, url, err)
			os.Exit(1)
		}
		fmt.Printf("\nfetch: statusCode:%d\n", resp.StatusCode)
	}
}

func hasValidPrefix(url string) bool {
	httpPrefixes := []string{"http://", "https://"}
	for _, prefix := range httpPrefixes {
		if strings.HasPrefix(url, prefix) {
			return true
		}
	}
	return false
}

func test() {
	v := true
	if v {
		fmt.Println("salut")
	}
}
