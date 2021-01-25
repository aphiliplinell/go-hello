package hello

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(server *Server) *mux.Router {
	router := mux.NewRouter()

	router.Use(loggingMiddleware)
	router.HandleFunc("/hello", server.SayHello).Methods("GET")
	router.HandleFunc("/redirect", server.OpenRedirect).Methods("GET")

	return router
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
