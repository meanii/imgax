package server

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/meanii/imgax/config"
)

type server struct {
	Port       string
	Host       string
	App        *fiber.App
	RootRouter fiber.Router
}

var Server *server

func (s *server) newServer() *server {
	app := fiber.New(fiber.Config{
		CaseSensitive: true,
		StrictRouting: false,
		Prefork:       true,
		ServerHeader:  "MAI",
		AppName:       "MAI v0.1-dev",
		JSONEncoder:   json.Marshal,
		JSONDecoder:   json.Unmarshal,
	})

	app.Use(cors.New())
	app.Use(compress.New())
	app.Use(requestid.New(requestid.Config{
		Header: "X-MAI-Request-ID",
	}))
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${url}​\n",
	}))

	s.App = app
	return s
}

func (s *server) StartServer() {
	addr := fmt.Sprintf("%s:%s", s.Host, s.Port)
	log.Infof("Server started at http://%s", addr)
	s.App.Listen(addr)
}

func InitServer() *server {
	s := server{
		Port: config.Env.Port,
		Host: config.Env.Host,
	}
	s.newServer()
	Server = &s
	return Server
}
