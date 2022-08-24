package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(url1, url2 string) (winner string, error error) {
	// aDuration := measureResponse(url1)
	// bDuration := measureResponse(url2)
	//
	// if aDuration < bDuration {
	//   return url1
	// }
	// return url2

	return ConfigurableRacer(url1, url2, tenSecondTimeout)
}

func ConfigurableRacer(url1, url2 string, timeout time.Duration) (winner string, error error) {
	// aDuration := measureResponse(url1)
	// bDuration := measureResponse(url2)
	//
	// if aDuration < bDuration {
	//   return url1
	// }
	// return url2

	select {
	case <-ping(url1):
		return url1, nil
	case <-ping(url2):
		return url2, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", url1, url2)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func measureResponse(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
