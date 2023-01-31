package main

import (
    "github.com/gofiber/fiber/v2"
	"os"
	"log"
	"api/internal/handlers"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func index(c *fiber.Ctx) error {
	return c.SendString("Hello from go server 2!")
}

func main() {
	os.Setenv("PORT", ":8080")

	port := os.Getenv("PORT")

	app := fiber.New() 

	app.Get("/", index) 

	app.Post("/product", handlers.CreateProduct)

	log.Fatal(app.Listen(port))
}