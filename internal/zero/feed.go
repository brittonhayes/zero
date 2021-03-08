package zero

import (
	"context"
	"sync"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Provider is an individual
// configuration consumed
// by the RSS parser
type Provider struct {
	Name    string `json:"name"`
	URL     string `json:"url"`
	Depth   int    `json:"depth"`
	Pattern string `json:"pattern"`
}

// Config is the parent
// RSS feed configuration
//
// It is used to set global
// and item specific parsing
// parameters
type Config struct {
	Global Global
	Items  []Provider
}

// Global is the universal
// set of configs that are
// applied to all feeds
type Global struct {
	Patterns []string
}

// TODO pull from channel of stories and start looking for regex matches
// TODO return matches as a list of results

// Setup creates a new instance
// of Config and reads in the user's
// config
func Setup() *Config {
	f := new(Config)
	if err := viper.UnmarshalKey("feeds", &f); err != nil {
		panic(err)
	}
	return f
}

// ReadRSS reaches out to feed sources
// and checks for pattern matches
func (c *Config) ReadRSS() Jobs {
	var wg sync.WaitGroup
	var results []Job
	jobs := make(chan Job, len(c.Items))

	// Send API responses to jobs channel
	go c.requestFeeds(jobs)

	wg.Add(1)
	go func() {
		// Wait for all items in
		// config
		for j := range jobs {
			results = append(results, j)
		}
		wg.Done()
	}()

	wg.Wait()
	return results
}

// * Pull from jobs channel and send to results
func (c *Config) requestFeeds(jobs chan<- Job) {
	var wg sync.WaitGroup

	// Populate sender channel
	// of stories
	client := gofeed.NewParser()

	// Range over items in config
	for _, item := range c.Items {
		wg.Add(1)
		go func(item Provider) {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			// Request entire feed from API
			// endpoint for each source
			response, err := client.ParseURLWithContext(item.URL, ctx)
			if err != nil {
				logrus.Errorf("error with URL: %s (%s)", item.URL, err.Error())
				defer wg.Done()
				return
			}

			logrus.WithFields(logrus.Fields{
				"STATUS":   "Requesting",
				"PROVIDER": item.Name,
				"PATTERN":  item.Pattern,
			}).Debug()

			// Write API response to channel of jobs
			j := NewJob(&item, response)
			jobs <- j

			// Set goroutine as
			// done
			wg.Done()
		}(item)
	}

	// Wait for goroutine
	// to complete
	wg.Wait()
	close(jobs)
}
