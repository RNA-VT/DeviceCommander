package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/graph"
	"github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/postgres"
)

func (a *APIService) addGraphQLRoutes(e *echo.Echo, deviceService *postgres.DeviceService) {
	baseRoute := "/v1/graphql"
	api := e.Group(baseRoute)

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				DeviceService: deviceService,
			}},
		),
	)

	api.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	playgroundHandler := playground.Handler("GraphQL", baseRoute+"/query")
	api.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})
}
