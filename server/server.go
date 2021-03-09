package server

import (
	"net/http"
	"sync"

	log "github.com/sirupsen/logrus"
)

func Run() {
	mux := http.NewServeMux()

	// Setup wait group for server
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Initialize handlers
	mux.Handle("/", middleware(http.HandlerFunc(dashboard)))

	// Goroutine for webserver
	log.Info("HTTP Server Started")
	go func() {
		log.Fatal(http.ListenAndServe(":8091", mux))
		wg.Done()
	}()

	// Wait until done
	// to close function
	wg.Wait()
}
