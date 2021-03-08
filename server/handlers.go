package server

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func getMatches(w http.ResponseWriter, req *http.Request) {
	log.Infof(LogInfof, req.Method, req.URL)

	// f := intel.New()
	// Setup channel of feed
	// items
	// var matches intel.Matches
	// results := make(chan intel.Item)
	// f.Setup(results)
	// f.ReadRSS(results, w)
}
