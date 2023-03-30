package controller

import (
	"net/http"
)

func CheckUrls(urls []string) bool {
	ch := make(chan bool, len(urls))
	for _, url := range urls {
		go func(url string) {
			_, err := http.Get(url)
			if err != nil {
				ch <- false
				return
			}
			ch <- true
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		if !<-ch {
			return false
		}
	}
	return true
}
