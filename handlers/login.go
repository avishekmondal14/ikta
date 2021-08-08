package handlers

import (
	"net/http"
	"os"
)

func AdminLogin(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(os.Getenv("host")))
}

func TeacherLogin(w http.ResponseWriter, req *http.Request) {

}
