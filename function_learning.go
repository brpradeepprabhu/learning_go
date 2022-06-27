package main

import (
	"fmt"
	"net/http"
)

func main() {
	content, err := contentType("https://linkedin.com")
	if err != nil {
		fmt.Println("Error is %s", err)
	} else {
		fmt.Println(content)
	}
}

func contentType(req string) (string, error) {
	resp, err := http.Get(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content := resp.Header.Get("Content-Type")
	fmt.Println("content %s", content)
	return content, nil

}
