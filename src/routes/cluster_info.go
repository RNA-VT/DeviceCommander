package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func (a *APIService) addInfoRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.GET("/cluster_info", a.getClusterInfo)
	api.GET("/microcontroller", a.getMicrocontrollers)
	api.GET("/microcontroller/:id", a.getMicrocontroller)
	api.GET("/component", a.getComponents)
	api.GET("/component/:id", a.getComponent)
	api.GET("/config", a.getComponentConfig)
}

func (a APIService) getClusterInfo(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster.GetConfig())
}

func (a APIService) getMicrocontrollers(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster.Microcontrollers)
}

func (a APIService) getMicrocontroller(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusTeapot, "Fuck")
	}
	micros := a.Cluster.GetMicrocontrollers()
	micro, ok := micros[id]
	if !ok {
		return c.JSON(http.StatusTeapot, "Fuck")
	}
	return c.JSON(http.StatusOK, micro)
}

func (a APIService) getComponents(c echo.Context) error {
	return c.JSON(http.StatusOK, a.Cluster.Me.GetComponentMap())
}

func (a APIService) getComponent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	for i := 0; i < len(a.Cluster.Me.Solenoids); i++ {
		if a.Cluster.Me.Solenoids[i].UID == id {
			return c.JSON(http.StatusOK, a.Cluster.Me.Solenoids[i])
		}
	}
	return c.JSON(http.StatusNotFound, "ID Not Found")
}

func (a APIService) getComponentConfig(c echo.Context) error {
	yamlFile, err := ioutil.ReadFile("./app/config/microcontroller.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return err
	}
	return c.JSON(http.StatusOK, string(yamlFile))
}
