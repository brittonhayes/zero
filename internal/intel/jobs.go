package intel

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mmcdole/gofeed"
)

type Job struct {
	Provider *Provider    `json:"provider"`
	Feed     *gofeed.Feed `json:"feed"`
}

type Jobs []Job

func NewJob(feed *Provider, response *gofeed.Feed) Job {
	return Job{Provider: feed, Feed: response}
}

func (jobs Jobs) FindMatches() ([]*Match, error) {
	var matches []*Match
	for _, j := range jobs {
		m := findMatch(j)
		if m != nil {
			matches = m
		}
	}

	return matches, nil
}

func (jobs Jobs) List() []*gofeed.Feed {
	var feeds []*gofeed.Feed
	for _, j := range jobs {
		feeds = append(feeds, j.Feed)
	}

	return feeds
}

func findMatch(j Job) []*Match {
	var matches []*Match
	reg, err := regexp.Compile(j.Provider.Pattern)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(j.Feed.Items); i++ {
		if i <= j.Provider.Depth {
			fields := fmt.Sprintf("%s %s %s", j.Feed.Items[i].Description, j.Feed.Items[i].Content, strings.Join(j.Feed.Items[i].Categories, ""))
			if reg.MatchString(fields) {
				matches = append(matches, NewMatch(j.Provider, j.Feed, j.Feed.Items[i]))
			}
		}
	}

	return matches
}
