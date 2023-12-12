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

func (f *FiberServer) Get(path string, handlers ...fiber.Handler) {
	f.server.Get(path, handlers...)
}

func (f *FiberServer) Group(prefix string) fiber.Router {
	return f.server.Group(prefix)
}

func (f *FiberServer) Start() error {
	return f.server.Listen(":3000")
}
