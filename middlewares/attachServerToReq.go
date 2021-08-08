package middlewares

import (
	"context"
	"net/http"

	"github.com/avishekmondal14/ikta/server"
)

// Server - Embeds the main server for use in middlewares
type Server struct {
	*server.Server
}

// AttachServerToReq - Middleware attaches Server instance to the request in its context
func (server *Server) AttachServerToReq(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.WithValue(req.Context(), "server", server.Server)

		// next.ServeHTTP(w, req.WithContext(ctx))
		next.ServeHTTP(w, req.Clone(ctx))
	})
}
