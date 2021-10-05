package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a APIService) addManageRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	// api.POST("/:id", a.editDevice)
	api.POST("/component/:id", a.editComponent)
}

func (a APIService) editComponent(c echo.Context) error {
	log.Println("start editing component")

	body := c.Request().Body

	wholeBody, err := ioutil.ReadAll(body)
	if err != nil {
		return echo.NewHTTPError(
			http.StatusNotAcceptable,
			"Please provide valid Request Body")
	}

	message := map[string]interface{}{}
	err = json.Unmarshal([]byte(wholeBody), &message)
	if err != nil {
		return err
	}

	// component, err := a.Cluster.GetComponent(c.Param("id"))

	// if err != nil {
	// 	return echo.NewHTTPError(
	// 		http.StatusNotAcceptable,
	// 		"Error getting component")
	// }

	for key, value := range message {
		fmt.Println(key, value)
		if key == "Enabled" {
			fmt.Println("CHANGE enabled")
			fmt.Println(c.Param("id"))
			// GET COMPONENT AND MAKE EDIT
			// a.Cluster.Me.Description = value.(string)
		}
	}

	return c.JSON(http.StatusOK, c.Param("id"))
}

// func (a APIService) editDevice(c echo.Context) error {
// 	log.Println("start editing device")

// 	body := c.Request().Body

// 	wholeBody, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		return echo.NewHTTPError(
// 			http.StatusNotAcceptable,
// 			"Please provide valid Request Body")
// 	}

// 	message := map[string]interface{}{}
// 	err = json.Unmarshal([]byte(wholeBody), &message)
// 	if err != nil {
// 		return err
// 	}

// 	for key, value := range message {
// 		fmt.Println(key, value)
// 		if key == "description" {
// 			fmt.Println("CHANGE DESCRIPTION")
// 			// a.Cluster.Me.Description = value.(string)
// 		}
// 	}

// 	return c.JSON(http.StatusOK, c.Param("id"))
// }
