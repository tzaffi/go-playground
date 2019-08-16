package main

import (
	"fmt"
	"sync"
)

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Zero sets the counter to 0 for the given key.
func (c *SafeCounter) Zero(key string) {
	c.mux.Lock()
	c.v[key] = 0
	c.mux.Unlock()
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) (int, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.v[key]
	return val, ok
}

func (c *SafeCounter) String() string {
	s := ""
	for k, v := range c.v {
		s += fmt.Sprintf("%v->%v\n", k, v)
	}
	return s
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlButWrong(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func Crawl(url string, depth int, fetcher Fetcher) {
	urlAttempts := SafeCounter{v: make(map[string]int)}

	var wg = &sync.WaitGroup{}
	var crawl func(url string, depth int)
	crawl = func(url string, depth int) {
		defer wg.Done()
		if depth <= 0 {
			return
		}
		cnt, ok := urlAttempts.Value(url)
		if !ok {
			body, urls, err := fetcher.Fetch(url)
			if err != nil {
				urlAttempts.Zero(url)
				fmt.Println(err)
				return
			}
			urlAttempts.Inc(url)

			fmt.Printf("found: %s %q\n", url, body)
			for _, u := range urls {
				wg.Add(1)
				go crawl(u, depth-1)
			}
			return
		}
		if cnt > 0 {
			urlAttempts.Inc(url)
		}
		fmt.Printf("not trying: %s\n", url)

	}
	wg.Add(1)
	crawl(url, depth)
	wg.Wait()
	fmt.Printf("urlAttempts = %v\n", &urlAttempts)
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
