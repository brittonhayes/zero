package server

import (
	"net/http"
	"sync"
	"time"

	"github.com/hashicorp/go-uuid"
	log "github.com/sirupsen/logrus"
)

type Response struct {
	ID     string      `json:"id"`
	Time   string      `json:"time"`
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func NewResponse(status int, data interface{}) *Response {
	id, _ := uuid.GenerateUUID()
	return &Response{ID: id, Time: time.Now().Format(time.RFC3339), Status: status, Data: data}
}

func Run() {
	// Setup wait group for server
	wg := new(sync.WaitGroup)
	wg.Add(1)

	// Initialize handlers
	http.HandleFunc("/", getAll)
	http.HandleFunc("/matches", getMatches)

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
