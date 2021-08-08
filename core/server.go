package core

import (
	"net/http"
	"os"

	"github.com/avishekmondal14/ikta/server"
)

func NewServer() (*server.Server, error) {
	server := &server.Server{}

	db, err := newDB()
	if err != nil {
		return nil, err
	}

	server.DB = db

	server.HTTPServer = &http.Server{}
	server.HTTPServer.Addr = os.Getenv("host") + ":" + os.Getenv("port")
	server.HTTPServer.Handler = GetHandler(NewRouter(server))

	return server, nil
}

func StartServer(server *http.Server) error {
	return server.ListenAndServe()
}

func ShutdownServer(server *server.Server) {
	// TODO
}
