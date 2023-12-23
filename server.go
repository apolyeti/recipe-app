package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/gocolly/colly"

	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.Debugger(&debug.LogDebugger{}))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// post request, get url from body and scrape it for ingredients, using regex
	// regex will look for numbers, and measurement words, such as cups, teaspoons, etc.
	// return a json object with the ingredients
	e.POST("/ingredients", func(c echo.Context) error {
		url := c.FormValue("url")
		c.Bind(&url)
		c.JSON(http.StatusOK, url)
		return nil
	})

	e.Logger.Fatal(e.Start(":1323"))
}
