package clienthandler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"siteavliable/internal/metrics"
	"siteavliable/internal/models"
)

type IClientCase interface {
	GetWithMinResponeTime(context.Context) (models.AccessTime, error)
	GetWithMaxResponeTime(context.Context) (models.AccessTime, error)
	GetByURL(context.Context, string) (models.AccessTime, error)
}

type HandlerRoutes struct {
	uCase  IClientCase
	logger *log.Logger
}

// NewHandlerRoutes returns a new instance HandlerRoutes
func NewHandlerRoutes(s IClientCase, l *log.Logger) *HandlerRoutes {
	return &HandlerRoutes{uCase: s, logger: l}
}

// GetMinAccessTime returns a new http.handler
func (h *HandlerRoutes) GetMinAccessTime() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		metrics.IncCounter(metrics.GetWithMinTime)
		accessTime, err := h.uCase.GetWithMinResponeTime(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Printf("Handler GetMinAccessTime Error: %s\n", err.Error())
			return
		}
		msg, _ := json.Marshal(accessTime)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

// GetMaxAccessTime returns a new http.handler
func (h *HandlerRoutes) GetMaxAccessTime() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		metrics.IncCounter(metrics.GetWithMaxTime)
		accessTime, err := h.uCase.GetWithMaxResponeTime(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Printf("Handler GetMaxAccessTime Error: %s\n", err.Error())
			return
		}
		msg, _ := json.Marshal(accessTime)

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

// GetByURL returns a new http.handler
func (h *HandlerRoutes) GetByURL() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		metrics.IncCounter(metrics.GetByURL)
		params := r.URL.Query().Get("url")
		if params == "" {
			http.Error(w, "Bag query params", http.StatusBadRequest)
			return
		}
		accessTime, err := h.uCase.GetByURL(r.Context(), params)
		if err != nil {
			h.logger.Printf("Handler GetByUrl Error: %s\n", err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		msg, err := json.Marshal(accessTime)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

// ClientsRouter registers endpoints in router
func ClientsRouter(mux *http.ServeMux, s IClientCase, l *log.Logger) {
	//config router
	h := &HandlerRoutes{uCase: s, logger: l}
	mux.Handle("/min", h.GetMinAccessTime())
	mux.Handle("/max", h.GetMaxAccessTime())
	mux.Handle("/url", h.GetByURL())
}
