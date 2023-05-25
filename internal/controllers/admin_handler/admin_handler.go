package adminhandler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"siteavliable/internal/models"
)

// IAdminCase interface
type IAdminCase interface {
	GetMetrics(ctx context.Context) ([]models.CounterStats, error)
}

// HandlerAdmins...
type HandlerAdmins struct {
	uCase    IAdminCase
	logger   *log.Logger
	authCred models.AuthCred
}

// GetMetrics returns a new htp.Handler
func (h *HandlerAdmins) GetMetrics() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		metrics, err := h.uCase.GetMetrics(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Printf("Handler GetMetrics error: %s\n", err.Error())
			return
		}

		msg, err := json.Marshal(metrics)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			h.logger.Printf("Handler GetMetrics marshal error: %s\n", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
	return http.HandlerFunc(fn)
}

// AdminRouter registers path new enpoint in router
func AdminRouter(mux *http.ServeMux, c IAdminCase, l *log.Logger, auth models.AuthCred) {
	h := &HandlerAdmins{uCase: c, logger: l, authCred: auth}
	handler := h.AdminAuthMiddleware(h.GetMetrics())
	mux.Handle("/metrics", handler)
}

func (h *HandlerAdmins) AdminAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok || user != h.authCred.User || pass != h.authCred.Pass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Unauthorized access")
			return
		}

		// Pass the request to the next handler
		next.ServeHTTP(w, r)
	})
}
