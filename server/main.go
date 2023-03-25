package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santacodes/SecureEx/server/api"
)

type jobj struct {
	Domain       string `json:"domain"`
	Authenticity int    `json:authenticity`
	Safety       int    `json:safety`
}

func main() {
	fmt.Println("\n")
	fmt.Println("------ SECUREX --------")
	fmt.Println("\n")
	fmt.Println("Fetching results from API")
	fmt.Println("\n")
	//pass the domain here
	api.GetInfo("gsocorganizations.dev")
	fmt.Println("\n")
	fmt.Println("Hosting the Server!")
	//server init

	p1 := jobj{
		Domain:       "google.com",
		Authenticity: 50,
		Safety:       60,
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the SecureEX API. To access the API go to ip/\"domain\"")
	})

	jsonBytes, err := json.Marshal(p1)
	fmt.Println(string(jsonBytes), err)
	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			//call the api.go here and get the details of the website
			return c.SendString(string(jsonBytes))
			// => JSON String of struct jobj
		}
		return c.SendString("No Domain Name provided")
	})
	app.Listen(":3000")

}
