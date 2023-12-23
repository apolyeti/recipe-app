package main

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
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
	Name        string `json:"ingredient"`  // name of ingredient
	Quantity    string `json:"quantity"`    // quantity of ingredient
	Measurement string `json:"measurement"` // measurement of ingredient
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
				Name:        url,
				Quantity:    "0",
				Measurement: "cups",
			},
			{
				Name:        "ingredient2",
				Quantity:    "quantity2",
				Measurement: "grams",
			},
		},
	}
	scrapeURL(url)
	return c.JSON(http.StatusOK, r)
}

func scrapeURL(url string) {
	fmt.Println("scraping url:", url)
	// Instantiate default collector
	c := colly.NewCollector()

	ingredients := ""
	// find an element with a class that contains "ingredient"
	c.OnHTML("[class*=ingredient]", func(e *colly.HTMLElement) {
		ingredients += e.Text + "\n"
		fmt.Println(ingredients)
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Now use OpenAI Api to format the ingredients

}

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	// e.GET("/", func(c echo.Context) error {
	// 	return c.String(http.StatusOK, "Hello, World!")
	// })

	e.POST("/api/ingredients", getIngredients)

	// request body:
	// method: POST
	// headers: Content-Type: application/json
	// body: {"url": url}
	// for now, response body is the same as request body

	e.Logger.Fatal(e.Start(":1323"))
}
