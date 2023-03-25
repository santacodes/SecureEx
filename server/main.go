package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/santacodes/SecureEx/server/api"
)

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
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")

}
