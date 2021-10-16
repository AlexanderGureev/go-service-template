package apiserver

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type ApiServer struct{}

type response struct {
	Success bool
}

func Create(ctx context.Context) *ApiServer {
	return &ApiServer{}
}

func (s *ApiServer) Listen(port string) error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(response{Success: true})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(res)
	})

	server := &http.Server{
		Addr:           ":" + port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return server.ListenAndServe()
}
