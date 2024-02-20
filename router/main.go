package router

import (
	"github.com/gofiber/fiber/v2"

	"github.com/meanii/imgax/controller"
	"github.com/meanii/imgax/server"
)

type Router struct {
	app  *fiber.App
	root fiber.Router
}

func (r *Router) Init() {
	r.app = server.Server.App
	r.root = r.rootRouter()
	r.loadRouters()
}

func (r *Router) rootRouter() (root fiber.Router) {
	rootGroup := r.app.Group("/mai")
	server.Server.RootRouter = rootGroup
	return rootGroup
}

func (r *Router) loadRouters() {
	r.root.Get("/", controller.Home)
	r.root.Get("/pr", controller.Processor)
}
