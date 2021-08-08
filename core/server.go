package core

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/avishekmondal14/ikta/server"
)

func NewServer() (*server.Server, error) {
	server := &server.Server{}

	db, err := newDB()
	if err != nil {
		return nil, err
	}

	server.DB = db

	// os.Setenv("host", "localhost")
	// os.Setenv("port", "4000")

	server.HTTPServer = &http.Server{}
	server.HTTPServer.Addr = os.Getenv("host") + ":" + os.Getenv("port")
	server.HTTPServer.Handler = GetHandler(NewRouter(server))

	return server, nil
}

func StartServer(server *http.Server) error {
	return server.ListenAndServe()
}

func ShutdownServer(server *server.Server) {
	// Create channel to receive OS signal
	c := make(chan os.Signal, 1)

	// SIGINT (Ctrl+C) - Graceful shutdown
	// SIGQUIT (Ctrl+\) - Quit process without graceful shutdown
	// SIGKILL or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT)

	// Block until we receive our signal.
	sig := <-c

	if sig == syscall.SIGINT {
		// os.Setenv("graceful_shutdown_timeout", "15")
		gracefulShutdownTimeout, _ := strconv.ParseInt(os.Getenv("graceful_shutdown_timeout"), 10, 64)
		// Create a deadline to wait for.
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(rand.Int63n(gracefulShutdownTimeout)))
		defer cancel()
		// Doesn't block if no connections, but will otherwise wait
		// until the timeout deadline.

		// Close database connection
		sqlDB, err := server.DB.DB()
		if err != nil {
			log.Println("Error while trying to retrieve sql db from gorm. Quiting process.")
			log.Println(err)
			os.Exit(1)
		}
		err = sqlDB.Close()
		if err != nil {
			log.Println("Error while trying to close database connection. Quiting process.")
			log.Println(err)
			os.Exit(1)
		}

		// Shutdown server gracefully
		err = server.HTTPServer.Shutdown(ctx)
		if err != nil {
			log.Println("Error while trying to shutdown server gracefully. Quiting process.")
			log.Println(err)
			os.Exit(1)
		}
		// Optionally, you could run server.httpServer.Shutdown in a goroutine and block on
		// <-ctx.Done() if your application should wait for other services
		// to finalize based on context cancellation.

		fmt.Println()
		log.Println("Shutting down server gracefully")
		os.Exit(0)
	} else if sig == syscall.SIGQUIT {
		log.Println("Quiting process. No graceful shutdown.")
		os.Exit(1)
	}
}
