package concurrency

type WebsiteChecker func(string) bool

// As we don't need either value to be named, each of them is anonymous within the struct;
// this can be useful in when it's hard to know what to name a value.
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultChannel := make(chan result)

	// Because the only way to start a goroutine is to put go in front of a function call,
	// we often use anonymous functions when we want to start a goroutine.
	for _, url := range urls {
		// they can be executed at the same time that they're declared
		// this is what the () at the end of the anonymous function is doing
		// Secondly they maintain access to the lexical scope in which they are defined.
		go func(u string) {
			// But each of our goroutines have a reference to the url variable - they don't have their own independent copy.
			// results[u] = wc(u)
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
