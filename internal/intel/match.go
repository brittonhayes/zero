package intel

import "github.com/mmcdole/gofeed"

type Match struct {
	Provider *Provider
	Feed     *gofeed.Feed
	Item     *gofeed.Item
}

type Matches []Match

func NewMatch(provider *Provider, feed *gofeed.Feed, item *gofeed.Item) *Match {
	return &Match{Provider: provider, Feed: feed, Item: item}
}
