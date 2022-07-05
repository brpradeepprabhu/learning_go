package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	urlList := []string{
		"https://jsonplaceholder.typicode.com/todos/1",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}
	var startTime = time.Now()
	siteSerials(urlList)
	fmt.Println(time.Since(startTime))

	startTime = time.Now()
	siteConcurrently(urlList)
	fmt.Println(time.Since(startTime))

	startTime = time.Now()
	siteUsingChannel(urlList)
	fmt.Println(time.Since(startTime))

	//siteChannel()
}
func siteSerials(url []string) {
	for _, value := range url {
		resp, err := getContentType(value)
		if err == nil {
			fmt.Println("resp", resp)
		} else {
			fmt.Println("err", err)
		}
	}
}

func siteConcurrently(url []string) {
	var wg sync.WaitGroup
	for _, value := range url {
		wg.Add(1)
		go func(url string) {
			resp, err := getContentType(value)
			if err == nil {
				fmt.Println(" siteConcurrently resp", resp)
			} else {
				fmt.Println("err", err)
			}
			wg.Done()
		}(value)
	}
	wg.Wait()
}

func siteChannel() {
	ch := make(chan int)

	go func() {
		ch <- 32
	}()
	val := <-ch
	fmt.Println("val", val)
	count := 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Println("sending message", i)
			ch <- i
			time.Sleep(1000)
		}
	}()

	for i := 0; i < count; i++ {
		val = <-ch
		fmt.Println("recieivng message", i)

	}

}
func getContentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", fmt.Errorf("error to get the content type %s", err)
	}
	resp.Body.Close()
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		return "", fmt.Errorf("error to get the content type")
	}
	return contentType, nil
}

func siteUsingChannel(url []string) {
	chanel := make(chan string)
	for _, value := range url {
		go getContentTypeUsingChanel(value, chanel)

	}
	for range url {
		output := <-chanel
		fmt.Println("Running from channel", output)
	}
}

func getContentTypeUsingChanel(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("Error in getting content type")
	}
	contentType := resp.Header.Get("Content-Type")
	if contentType == "" {
		out <- fmt.Sprintf("Error in getting content type")
	}
	out <- contentType
}
