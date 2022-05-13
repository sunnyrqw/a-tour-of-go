package tour

import (
	"testing"
	"time"
)

func TestWebCrawler(t *testing.T) {
	go Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second)
}
