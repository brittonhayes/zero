package intel

import (
	"fmt"
	"io"
	"regexp"
	"sync"
	"text/template"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// FeedItem is an individual
// configuration consumed
// by the RSS parser
type FeedItem struct {
	Name    string
	URL     string
	Depth   int
	Pattern string
}

// Feeds is the parent
// RSS feed configuration
//
// It is used to set global
// and item specific parsing
// parameters
type Feeds struct {
	Global Global
	Items  []FeedItem
}

// Global is the universal
// set of configs that are
// applied to all feeds
type Global struct {
	Patterns []string
}

// Match contains all of the
// fields required to check
// and print matched patterns
type Match struct {
	Pattern     *regexp.Regexp
	Reference   string
	ExtraFields []string
}

// New creates a new instance
// of Feeds and reads in the user's
// config
func New() Feeds {
	f := new(Feeds)
	if err := viper.UnmarshalKey("feeds", &f); err != nil {
		panic(err)
	}
	return *f
}

// TODO add method that checks for most commonly referenced topics

// Read reaches out to feed sources
// and checks for pattern matches
//
// It will then write out any matches found
// to the provided io.Writer
func (f Feeds) Read(items chan FeedItem, w io.Writer) {
	var wg sync.WaitGroup
	for item := range items {
		go item.fetch(&wg, w, f.Global.Patterns)
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

// fetch reaches out to an rss feed for details on an item
//
// It then checks the articles for pattern matches from
// the user-defined config
func (f FeedItem) fetch(wg *sync.WaitGroup, w io.Writer, keywords []string) {
	wg.Add(1)
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(f.URL)
	if err != nil {
		logrus.Warnf("Unable to read %s, %v", f.URL, err)
		wg.Done()
		return
	}

	// Global patterns
	go func() {
		wg.Add(1)
		for _, k := range keywords {
			global := regexp.MustCompile(k)
			for i := 0; i <= f.Depth; i++ {
				m := newMatch(global, feed.Items[i].Content, feed.Items[i].Link)
				m.print(w)
			}
		}
		wg.Done()
	}()

	// Per-Item patterns
	go func() {
		wg.Add(1)
		pattern := regexp.MustCompile(f.Pattern)
		for i := 0; i <= f.Depth; i++ {
			m := newMatch(pattern, feed.Items[i].Content, feed.Items[i].Link)
			m.print(w)
		}
		wg.Done()
	}()
	wg.Done()
}

// newMatch creates a new instance of the Match type
// for printing matching RSS feed results
func newMatch(pattern *regexp.Regexp, reference string, extraFields ...string) *Match {
	return &Match{Pattern: pattern, Reference: reference, ExtraFields: extraFields}
}

// print prints out the results of a match
// to an io.Writer
func (m *Match) print(w io.Writer) {
	if m.Pattern.MatchString(m.Reference) {
		_, _ = fmt.Fprintf(w, "Match: %q\n", m.ExtraFields)
	}
}

// render prints out the results of a match
// to an io.Writer using a go text/template
// for formatting
func (m *Match) render(w io.Writer, tmpl string) {
	if m.Pattern.MatchString(m.Reference) {
		t, err := template.New("match").Parse(tmpl)
		if err != nil {
			panic(err)
		}

		if err := t.ExecuteTemplate(w, "match", &m); err != nil {
			panic(err)
		}
	}
}
