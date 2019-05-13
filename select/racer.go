package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

// Racer compares the response times of a and b, returning the fastest one, timing out after 10s
func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// ConfigurableRacer compares the response times of a and b, returning the fastest one
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// "myVar := <-ch": This is a blocking call, as you're waiting for a value.
	// What "select" lets you do is wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
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
