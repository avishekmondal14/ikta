package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/avishekmondal14/ikta/core"
)

func main() {
	server, err := core.NewServer()
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		if err := core.StartServer(server.HTTPServer); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Fatalln(err)
			}
		}
	}()

	core.ShutdownServer(server)
}
