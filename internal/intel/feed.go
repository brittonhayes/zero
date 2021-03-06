package intel

import (
	"fmt"
	"io"
	"regexp"
	"sync"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type FeedItem struct {
	Name    string
	URL     string
	Depth   int
	Pattern string
}

type Feeds struct {
	Keywords []string
	Items    []FeedItem
}

func New() Feeds {
	f := new(Feeds)
	if err := viper.UnmarshalKey("feeds", &f); err != nil {
		panic(err)
	}
	return *f
}

// TODO add method that checks for most commonly referenced topics
func (f Feeds) Read(items chan FeedItem, w io.Writer) {
	var wg sync.WaitGroup
	for i, _ := range items {
		go fetchContent(f, i, &wg, w)
	}
	wg.Wait()
}

func (f Feeds) Setup(results chan<- FeedItem) {
	go func() {
		for _, item := range f.Items {
			results <- item
		}
		close(results)
	}()
}

func fetchContent(feeds Feeds, index int, wg *sync.WaitGroup, w io.Writer) {
	wg.Add(1)
	fp := gofeed.NewParser()
	item := feeds.Items[index]

	feed, err := fp.ParseURL(item.URL)
	if err != nil {
		logrus.Warnf("Unable to read %s, %v", item.URL, err)
		wg.Done()
		return
	}

	// Global patterns
	keywords := regexp.MustCompile("")
	// Per-Item patterns
	pattern := regexp.MustCompile(item.Pattern)

	for i := 0; i <= item.Depth; i++ {
		if pattern.MatchString(feed.Items[i].Content) || keywords.MatchString(feed.Items[i].Content) {
			_, _ = fmt.Fprintf(w, "%s\n", pattern.FindString(feed.Items[i].Content))
		}
	}
	wg.Done()
}
