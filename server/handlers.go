package server

import (
	"net/http"

	"github.com/brittonhayes/zero/internal/zero"
	log "github.com/sirupsen/logrus"
)

func LogRequest(r *http.Request) {
	log.WithFields(log.Fields{
		"METHOD": r.Method,
		"PATH":   r.URL,
		"HOST":   r.Host,
	}).Info()
}

func getMatches(w http.ResponseWriter, req *http.Request) {

	// Default request
	// logging
	LogRequest(req)

	matches, err := zero.Setup().ReadRSS().Inspect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	b, err := NewResponse(http.StatusOK, matches).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}

func getAll(w http.ResponseWriter, req *http.Request) {

	// Default request
	// logging
	defer LogRequest(req)

	z := zero.Setup()
	j := z.ReadRSS()

	b, err := NewResponse(http.StatusOK, j).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
}
