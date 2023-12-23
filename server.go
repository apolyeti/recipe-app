package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Ingredient struct {
	/*
		json: {ingredients: [
			{ingredient: "ingredient1", quantity: "quantity1"},
			{ingredient: "ingredient2", quantity: "quantity2"},
			...
		]}
	*/
	// should be list of ingredients
	Name     string `json:"ingredient"` // name of ingredient
	Quantity string `json:"quantity"`   // quantity of ingredient
}

type Recipe struct {
	/*
		json: {recipe: [
			{ingredient: "ingredient1", quantity: "quantity1"},
			{ingredient: "ingredient2", quantity: "quantity2"},
			...
		]}
	*/
	// should be list of ingredients
	Ingredients []Ingredient `json:"recipe"`
}

type URL struct {
	Name string `json:"url"`
}

func getIngredients(c echo.Context) error {
	// get url from request body
	requestBody := new(URL)

	if err := c.Bind(requestBody); err != nil {
		return err
	}

	// get url from req body
	// print url to console
	url := requestBody.Name
	fmt.Println("url:", url)
	r := &Recipe{
		Ingredients: []Ingredient{
			{
				Name:     url,
				Quantity: "0",
			},
			{
				Name:     "ingredient2",
				Quantity: "quantity2",
			},
		},
	}
	return c.JSON(http.StatusOK, r)
}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/api/ingredients", getIngredients)

	// request body:
	// method: POST
	// headers: Content-Type: application/json
	// body: {"url": url}
	// for now, response body is the same as request body

	e.Logger.Fatal(e.Start(":1323"))
}
