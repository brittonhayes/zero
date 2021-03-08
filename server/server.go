package server

import (
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
)

func Run() {
	// Setup wait group for server
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// Initialize handlers
	http.HandleFunc("/", getMatches)

	// Goroutine for API server
	go func() {
		log.Fatal(http.ListenAndServe(":8090", nil))
		wg.Done()
	}()

	// Goroutine for webserver
	log.Info("HTTP Server Started")
	go func() {
		log.Fatal(http.ListenAndServe(":8091", nil))
		wg.Done()
	}()

	// Wait until done
	// to close function
	wg.Wait()
}
