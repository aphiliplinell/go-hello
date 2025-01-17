package hello

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct{}

func (s *Server) SayHello(w http.ResponseWriter, r *http.Request) {
	response := struct {
		Message string `json:"message"`
	}{
		Message: "Hello",
	}

	data, err := json.Marshal(response)
	if err != nil {
		log.Printf("could not marshal data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (s *Server) OpenRedirect(w http.ResponseWriter, r *http.Request) {
	// https://codeql.github.com/codeql-query-help/go/go-unvalidated-url-redirection/
	r.ParseForm()
	http.Redirect(w, r, r.Form.Get("target"), http.StatusFound)
}
