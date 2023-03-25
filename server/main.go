package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/santacodes/SecureEx/server/api"
	"github.com/santacodes/SecureEx/server/database"
)

type jobj struct {
	Domain       string `json:"domain"`
	Authenticity int    `json:"authenticity"`
	Safety       int    `json:"safety"`
}

func main() {
	fmt.Println("------ SECUREX --------")
	database.DBMigrate()
	log.Println("Hosting the Server!")

	p1 := jobj{
		Domain:       "http://surprise.shopping/",
		Authenticity: 50,
		Safety:       60,
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is the SecureEX API. To access the API go to ip/\"domain\"")
	})

	app.Get("/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			//call the api.go here and get the details of the website
			api.GetInfo(c.Params("name"))
			p1.Domain = c.Params("name")
			jsonBytes, err := json.Marshal(p1)
			fmt.Println(string(jsonBytes), err)
			return c.SendString(string(jsonBytes))
			// => JSON String of struct jobj
		}
		return c.SendString("No Domain Name provided")
	})
	app.Listen(":3000")

}
