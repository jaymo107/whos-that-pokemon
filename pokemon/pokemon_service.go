package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type SpritesRes struct {
	FrontDefault string `json:"front_default"`
}

type PokemonRes struct {
	Name    string
	Sprites SpritesRes `json:"sprites"`
}

type Pokemon struct {
	Name  string
	Image string
}

type PokemonService struct {
	endpoint     string
	maxPokemonId int
}

func New() *PokemonService {
	return &PokemonService{
		endpoint:     "https://pokeapi.co/api/v2/pokemon/%d",
		maxPokemonId: 1025,
	}
}

// TODO:
// - Check there is an image there and the pokemon was found. Handle 404s
// - Store the pokemon when fetching in a local database to cache.
// - Keep track on pokemon that get correct responses the most.
// - Write some tests.
// - Make the templates better.

func (p *PokemonService) GetRandomPokemon() *Pokemon {
	id := rand.Intn(p.maxPokemonId)
	pokemon, _ := p.getPokemonById(id)

	return &Pokemon{
		Name:  pokemon.Name,
		Image: pokemon.Sprites.FrontDefault,
	}
}

func (p *PokemonService) getPokemonById(id int) (PokemonRes, error) {
	resp, err := http.Get(fmt.Sprintf(p.endpoint, id))

	fmt.Println(id)

	if err != nil {
		return PokemonRes{}, err
	}

	defer resp.Body.Close()

	pokemon := PokemonRes{}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return PokemonRes{}, err
	}

	json.Unmarshal(body, &pokemon)

	return pokemon, nil
}
