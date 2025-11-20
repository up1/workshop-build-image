package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// setupRoutes configures all the routes for the Echo server
func setupRoutes(e *echo.Echo) {
	e.GET("/", helloHandler)
}

// helloHandler handles the GET / endpoint
func helloHandler(c echo.Context) error {
	// Return json response
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
}

func main() {
	e := echo.New()
	setupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
