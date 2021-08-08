package handlers

import (
	"net/http"
)

// Options - Handler for request on any route with OPTIONS http method
func Options(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "No Content", http.StatusNoContent)
}
