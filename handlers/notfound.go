package handlers

import (
	"net/http"
)

// NotFound - Handler when no matching route is found
func NotFound(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not found"))
}
