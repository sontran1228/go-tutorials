package racer

import (
	"net/http"
	"time"
)

// Racer compares the response times of a and b, returning the fastest one
func Racer(a, b string) (winner string) {
	// "myVar := <-ch": This is a blocking call, as you're waiting for a value.
	// What "select" lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func measureResponseTime(url string) time.Duration {
	startA := time.Now()
	http.Get(url)
	return time.Since(startA)
}
