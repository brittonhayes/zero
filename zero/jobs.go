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
type Job struct {
	ID       string       `json:"id"`
	Provider Provider    `json:"provider"`
	Feed     *gofeed.Feed `json:"feed"`
}

//easyjson:json
type Jobs []Job

func NewJob(feed Provider, response *gofeed.Feed) Job {
	id, _ := uuid.GenerateUUID()
	return Job{ID: id, Provider: feed, Feed: response}
}

func (jobs Jobs) Inspect() (Matches, error) {
	var matches Matches
	for _, j := range jobs {
		log.WithFields(log.Fields{
			"STATUS":   "Searching",
			"PROVIDER": j.Provider.Name,
			"PATTERN":  j.Provider.Pattern,
		}).Debug()
		m := match(j)
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

func match(j Job) Matches {
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
