// Concurrently fetch multiple URLs using goroutine
package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func fetchURL(url string, c chan string) {
	start_ts := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
		os.Exit(1)
	}
	// body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	// if err != nil {
	// 	fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
	// }
	// fmt.Printf("%s", body)

	latency := time.Since(start_ts).Seconds()
	c <- fmt.Sprintf("%.2fs %s", latency, url)
}

// E.g., to test, run
//  go run urlfetch.go https://github.com https://google.com https://baidu.com
func main() {
	ch := make(chan string)

	fetchPerURL := 100

	start_ts := time.Now()
	for _, url := range os.Args[1:] {
		for i := 0; i < fetchPerURL; i++ {
			go fetchURL(url, ch)
		}
	}

	for i := 0; i < (len(os.Args)-1)*fetchPerURL; i++ {
		<-ch
		// fmt.Println(s)
	}
	latency := time.Since(start_ts).Seconds()

	fmt.Printf("Total runtime %fs\n", latency)
}
