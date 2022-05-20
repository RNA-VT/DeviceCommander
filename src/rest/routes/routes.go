package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

type Router interface {
	RegisterRoutes(*echo.Echo)
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
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

	e.GET("/swagger/*", echoSwagger.WrapHandler)
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
