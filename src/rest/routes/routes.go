package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

type Router interface{}

type BaseRouter struct{}

func (r BaseRouter) RegisterRoutes(e *echo.Echo) {
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS restricted
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

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
	DeviceRouter{}.RegisterRoutes(e)
}
