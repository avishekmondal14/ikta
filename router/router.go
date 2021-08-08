package router

import (
	"net/http"
	"regexp"
)

type Router struct {
	basepath        string
	middlewares     []func(http.Handler) http.Handler
	endpoints       []*endpoint
	subrouters      []*Router
	notfoundHandler func(http.ResponseWriter, *http.Request)
}

type endpoint struct {
	path    string
	method  string
	handler func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	return &Router{basepath: ``}
}

func (router *Router) NewSubrouter(subBasepath string) *Router {
	newRouter := &Router{basepath: router.basepath + subBasepath}

	for _, mw := range router.middlewares {
		newRouter.middlewares = append(newRouter.middlewares, mw)
	}

	router.subrouters = append(router.subrouters, newRouter)

	return newRouter
}

func (router *Router) Use(mw func(http.Handler) http.Handler) {
	router.middlewares = append(router.middlewares, mw)
}

func (router *Router) SetHandler(path string, method string, handler func(http.ResponseWriter, *http.Request)) {
	router.endpoints = append(router.endpoints, &endpoint{path, method, handler})
}

func (router *Router) SetNotfoundHandler(nf func(http.ResponseWriter, *http.Request)) {
	router.notfoundHandler = nf
}

func (router *Router) GetHandler() http.Handler {
	return http.HandlerFunc(router.handlePaths)
}

func (router *Router) handlePaths(w http.ResponseWriter, req *http.Request) {
	routerQueue := []*Router{router}

	for len(routerQueue) > 0 {
		for i := 0; i < len(routerQueue[0].endpoints); i++ {
			pth := routerQueue[0].basepath + routerQueue[0].endpoints[i].path
			if pth == `` {
				pth = `\/`
			}

			rgx := regexp.MustCompile(`^` + pth + `(\/)?$`)

			if rgx.MatchString(req.URL.Path) && routerQueue[0].endpoints[i].method == req.Method {
				routerQueue[0].execMw(http.HandlerFunc(routerQueue[0].endpoints[i].handler), len(routerQueue[0].middlewares)-1).ServeHTTP(w, req)
				return
			}
		}

		for i := 0; i < len(routerQueue[0].subrouters); i++ {
			routerQueue = append(routerQueue, routerQueue[0].subrouters[i])
		}

		routerQueue = routerQueue[1:]
	}

	if router.notfoundHandler != nil {
		router.notfoundHandler(w, req)
	}
}

func (router *Router) execMw(handler http.Handler, itr int) http.Handler {
	if itr < 0 {
		return handler
	}
	if itr == 0 {
		return router.middlewares[0](handler)
	}

	return router.execMw(router.middlewares[itr](handler), itr-1)
}
