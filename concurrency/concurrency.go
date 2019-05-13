package concurrency

// WebsiteChecker checks a url, returning a bool
type WebsiteChecker (func(string) bool)

type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map
// of urls to the result of checking each url with the WebsiteChecker function
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	// anonymous functions have a number of features:
	// Firstly, they can be executed at the same time that the're declared - this is what the () at the end of the anonymous function is doing.
	// Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}
