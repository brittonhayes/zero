package zero

import "github.com/mmcdole/gofeed"

//easyjson:json
type Match struct {
	Provider *Provider    `json:"provider"`
	Item     *gofeed.Item `json:"item"`
	RawMatch string       `json:"raw_match"`
}

//easyjson:json
type Matches []Match

func NewMatch(provider *Provider, item *gofeed.Item, finding string) Match {
	return Match{Provider: provider, Item: item, RawMatch: finding}
}
