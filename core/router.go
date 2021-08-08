package core

import (
	"net/http"

	"github.com/avishekmondal14/ikta/handlers"
	"github.com/avishekmondal14/ikta/middlewares"
	"github.com/avishekmondal14/ikta/router"
	"github.com/avishekmondal14/ikta/server"
)

func NewRouter(server *server.Server) *router.Router {
	rootRouter := router.NewRouter()
	rootRouter.Use((&middlewares.Server{server}).AttachServerToReq)
	rootRouter.SetNotfoundHandler(handlers.NotFound)
	rootRouter.SetHandler(`.*`, "OPTIONS", handlers.Options)

	loginRouter := rootRouter.NewSubrouter(``)
	loginRouter.SetHandler(`\/adminlogin`, "POST", handlers.AdminLogin)
	loginRouter.SetHandler(`\/teacherlogin`, "POST", handlers.TeacherLogin)

	return rootRouter
}

func GetHandler(router *router.Router) http.Handler {
	return router.GetHandler()
}
