package zero

import (
	"github.com/gammazero/workerpool"
	"github.com/mmcdole/gofeed"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Provider is an individual
// configuration consumed
// by the RSS parser
//easyjson:json
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

// NewConfig creates a new instance of the
// config object
func NewConfig(global Global, items []Provider) *Config {
	return &Config{Global: global, Items: items}
}

// Global is the universal
// set of configs that are
// applied to all feeds
type Global struct {
	Patterns []string
}

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
//
//
// It returns a slice of results that
// can be parsed for regex pattern
// matches.
func (c *Config) ReadRSS() Results {

	var results []Result
	jobs := make(chan Result, len(c.Items))

	// Send API responses to jobs channel
	c.requestFeeds(jobs)
	for j := range jobs {
		results = append(results, j)
	}

	return results
}

func (c *Config) requestFeeds(results chan<- Result) {
	defer close(results)

	wp := workerpool.New(4)
	// Populate sender channel
	// of stories
	client := gofeed.NewParser()

	// Range over items in config
	for _, item := range c.Items {
		item := item
		// Add requests to worker pool
		wp.Submit(func() {
			// Request entire feed from API
			// endpoint for each source
			response, err := client.ParseURL(item.URL)
			if err != nil {
				logrus.Errorf("error with URL: %s (%s)", item.URL, err.Error())
				return
			}

			logrus.WithFields(logrus.Fields{
				"STATUS":   "Requesting",
				"PROVIDER": item.Name,
				"PATTERN":  item.Pattern,
			}).Debug()

			// Write API response to channel of jobs
			result := NewResult(item, response)
			results <- result
		})
	}

	// Wait for worker pool
	// to finish
	wp.StopWait()
}
