package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Router interface {
	RegisterRoutes(*echo.Echo)
}

type BaseRouter struct {
	DeviceRouter
}

func (r BaseRouter) RegisterRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("/base", hello)

	r.registerFrontendRoutes(e)
	r.registerBackendRoutes(e)
}

func (r BaseRouter) registerFrontendRoutes(e *echo.Echo) {
	frontendRoot := "../frontend/build/"
	if viper.GetString("ENV") == "production" {
		frontendRoot = "/src/build/"
	}

	// Routes
	e.Static("/static", frontendRoot+"static")
	e.File("/*", frontendRoot+"index.html")
}

func (r BaseRouter) registerBackendRoutes(e *echo.Echo) {
	e.GET("/backend", hello)
	api := e.Group("group")
	api.GET("/base2", hello)
	r.DeviceRouter.RegisterRoutes(e)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
