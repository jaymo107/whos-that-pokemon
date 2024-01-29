package main

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/jaymo107/whos-that-pokemon/pokemon"
)

func main() {
	engine := html.New("./views", ".html")
	pokeService := pokemon.New()

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c fiber.Ctx) error {
		pokemon := pokeService.GetRandomPokemon()
		return c.Render("index", fiber.Map{
			"Name":  pokemon.Name,
			"Image": pokemon.Image,
		})
	})

	app.Post("/guess", func(c fiber.Ctx) error {
		pokemon := c.FormValue("pokemon")
		guess := c.FormValue("guess")

		template := "wrong"
		if strings.EqualFold(pokemon, guess) {
			template = "correct"
		}

		return c.Render(template, fiber.Map{
			"Name":  pokemon,
			"Image": c.FormValue("image"),
		})
	})

	app.Listen(":8080")
}
