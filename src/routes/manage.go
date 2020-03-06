package routes

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func (a APIService) addManageRoutes(e *echo.Echo) {
	api := e.Group("/v1")
	api.POST("/component/:id", a.editComponent)
}

func (a APIService) editComponent(c echo.Context) error {
	log.Println("start editing component")
	// body, err := c.Request().GetBody()
	// if err != nil {
	// 	log.Println("Failed to get warning message body")
	// }
	// fmt.Println(body)

	return c.JSON(http.StatusOK, c.Param("id"))

	// component, err := a.Cluster.GetComponent(c.Param("id"))
	// if err != nil {
	// 	return err
	// }

	// body := c.Request().Body
	// wholeBody, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	return echo.NewHTTPError(
	// 		http.StatusNotAcceptable,
	// 		"Please provide valid Request Body")
	// }

	// message := map[string]interface{}{}
	// err = json.Unmarshal([]byte(wholeBody), &message)
	// if err != nil {
	// 	return err
	// }

	// pin = c.Param("newPin")
	// newPin, err := component.EditPin(pin)

	// if err != nil {
	// 	return err
	// }

	// return c.JSON(http.StatusOK, component)
}
