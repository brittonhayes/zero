package server

import (
	"net/http"

	"github.com/brittonhayes/zero/templates"
	"github.com/brittonhayes/zero/zero"
	log "github.com/sirupsen/logrus"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"METHOD": r.Method,
			"PATH":   r.URL,
			"HOST":   r.Host,
		}).Info()
		next.ServeHTTP(w, r)
	})
}

func matches(w http.ResponseWriter, req *http.Request) {
	m, err := zero.Setup().ReadRSS().Inspect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	b, err := NewResponse(http.StatusOK, m).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func getAll(w http.ResponseWriter, req *http.Request) {
	z := zero.Setup()
	j := z.ReadRSS()

	b, err := NewResponse(http.StatusOK, j).MarshalJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	z := zero.Setup()
	results := z.ReadRSS()
	m, err := results.Inspect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Error(err)
		return
	}

	templates.Content(w, map[string]interface{}{
		"Matches": m,
		"Results": results,
	})
}
