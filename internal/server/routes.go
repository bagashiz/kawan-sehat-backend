package server

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.Handle("GET /", handle(notFound()))
	mux.Handle("GET /{$}", handle(index()))
}
