package routes

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"

	"github.com/rna-vt/devicecommander/graph"
	"github.com/rna-vt/devicecommander/graph/generated"
	"github.com/rna-vt/devicecommander/src/device"
	"github.com/rna-vt/devicecommander/src/endpoint"
)

func (a *APIService) addGraphQLRoutes(e *echo.Echo, deviceRepository device.Repository, endpointRepository endpoint.Repository) {
	baseRoute := "/v1/graphql"
	api := e.Group(baseRoute)

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{
				DeviceRepository:   deviceRepository,
				EndpointRepository: endpointRepository,
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
