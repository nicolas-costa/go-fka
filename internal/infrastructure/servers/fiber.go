package servers

import "github.com/gofiber/fiber/v2"

type Router interface {
	SetRoutes(app *FiberServer)
}

type FiberServer struct {
	server *fiber.App
}

func NewFiberServer(router Router) *FiberServer {
	app := fiber.New()

	server := &FiberServer{
		server: app,
	}

	router.SetRoutes(server)

	return server
}

func (f *FiberServer) Get(path string, handler ...fiber.Handler) {
	f.server.Get(path, handler...)
}

func (f *FiberServer) Start() error {
	return f.server.Listen(":3000")
}
