package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

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

// create struct for OpenAI response body
type OpenAI struct {
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		Logprobs      string `json:"logprobs"`
		Finish_reason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		Prompt_tokens     int `json:"prompt_tokens"`
		Completion_tokens int `json:"completion_tokens"`
		Total_tokens      int `json:"total_tokens"`
	} `json:"usage"`
	System_fingerprint string `json:"system_fingerprint"`
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
	c.OnHTML("[class*=ingredients]", func(e *colly.HTMLElement) {
		ingredients += e.Text + "\n"
		fmt.Println(ingredients)
	})

	err := c.Visit(url)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Now use OpenAI Api to format the ingredients
	recipe_list := openAI(ingredients)
	fmt.Println("recipe_list:\n", recipe_list)
}

func openAI(ingredients string) string {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY not set")
		os.Exit(1)
	}

	apiEndpoint := "https://api.openai.com/v1/chat/completions"

	promptFile := "prompt.txt"
	prompt, err := ioutil.ReadFile(promptFile)
	prompt_string := string(prompt)

	if err != nil {
		fmt.Println("error:", err)
	}

	payload := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "system", "content": prompt_string},
			{"role": "user", "content": ingredients},
		},
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("error:", err)
		return "ERROR"
	}

	req, err := http.NewRequest("POST", apiEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Println("error:", err)
		return "ERROR"
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error:", err)
		return "ERROR"
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error:", err)
		return "ERROR"
	}

	fmt.Println(string(body))

	var openAIResponse OpenAI
	err = json.Unmarshal(body, &openAIResponse)
	if err != nil {
		fmt.Println("error:", err)
		return "ERROR"
	}

	if len(openAIResponse.Choices) == 0 {
		fmt.Println("error: no choices")
		return "ERROR"
	}

	return openAIResponse.Choices[0].Message.Content

	// get body.choices.message.content[0]

	// return body.choices.message.content[0]
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
