package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/template/html/v2"
	"github.com/jaymo107/whos-that-pokemon/pokemon"
	"github.com/jaymo107/whos-that-pokemon/storage"
)

func main() {
	engine := html.New("./views", ".html")

	repository := storage.NewSqlRepository()

	pokeService := pokemon.NewPokemonService(
		pokemon.PokemonServiceConfig{
			Repository: repository,
		},
	)

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", func(c fiber.Ctx) error {
		pokemon := pokeService.GetRandomPokemon()
		return c.Render("index", fiber.Map{
			"Name":  pokemon.Name,
			"Image": pokemon.Image,
			"Id":    pokemon.Id,
		})
	})

	app.Get("/leaderboard", func(c fiber.Ctx) error {
		pokemon := repository.All()

		for i := range pokemon {
			pokemon[i].FormatName()
		}

		return c.Render("leaderboard", fiber.Map{
			"Pokemon": pokemon,
		})
	})

	app.Post("/guess", func(c fiber.Ctx) error {
		guess := c.FormValue("guess")
		id, _ := strconv.Atoi(c.FormValue("id"))

		pokemon, _ := repository.FindById(id)

		fmt.Println("guess: ", guess, pokemon.Name, id)

		template := "wrong"
		if strings.EqualFold(pokemon.Name, guess) {
			template = "correct"
			pokeService.MarkCorrectGuess(id)
		} else {
			pokeService.MarkIncorrectGuess(id)
		}

		return c.Render(template, fiber.Map{
			"Name":  pokemon.Name,
			"Image": pokemon.Image,
		})
	})

	app.Listen(":8080")
}
