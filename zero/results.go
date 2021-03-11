package zero

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/go-uuid"
	"github.com/mmcdole/gofeed"
	log "github.com/sirupsen/logrus"
)

//easyjson:json
type Result struct {
	ID       string       `json:"id"`
	Provider Provider     `json:"provider"`
	Feed     *gofeed.Feed `json:"feed"`
}

//easyjson:json
type Results []Result

func NewResult(feed Provider, response *gofeed.Feed) Result {
	id, _ := uuid.GenerateUUID()
	return Result{ID: id, Provider: feed, Feed: response}
}

func (r Results) FindMatches() (Matches, error) {
	var matches Matches
	for _, result := range r {
		log.WithFields(log.Fields{
			"STATUS":   "Searching",
			"PROVIDER": result.Provider.Name,
			"PATTERN":  result.Provider.Pattern,
		}).Debug()
		m := match(result)
		if m != nil {
			matches = m
		}
	}

	return matches, nil
}

func (r Results) List() []*gofeed.Feed {
	var feeds []*gofeed.Feed
	for _, j := range r {
		feeds = append(feeds, j.Feed)
	}

	return feeds
}

func match(j Result) Matches {
	var matches Matches
	reg, err := regexp.Compile("(?i)" + j.Provider.Pattern)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(j.Feed.Items); i++ {
		if i <= j.Provider.Depth {
			fields := fmt.Sprintf("%s %s %s", j.Feed.Items[i].Description, j.Feed.Items[i].Content, strings.Join(j.Feed.Items[i].Categories, ""))
			if reg.MatchString(fields) {
				matches = append(matches, NewMatch(&j.Provider, j.Feed.Items[i], reg.FindString(fields)))
			}
		}
	}

	log.WithFields(log.Fields{
		"STATUS":   "Done",
		"PROVIDER": j.Provider.Name,
		"PATTERN":  j.Provider.Pattern,
	}).Debug()

	return matches
}
