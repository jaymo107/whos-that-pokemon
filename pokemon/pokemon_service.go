package pokemon

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/jaymo107/whos-that-pokemon/storage"
)

type SpritesRes struct {
	FrontDefault string `json:"front_default"`
}

type PokemonRes struct {
	Name    string
	Sprites SpritesRes `json:"sprites"`
}

type Pokemon struct {
	Id    int
	Name  string
	Image string
}

type PokemonService struct {
	endpoint     string
	maxPokemonId int
	repository   storage.RepositoryInterface
	randFunc     RandFunc
}

type RandFunc func(int) int

type PokemonServiceConfig struct {
	Repository   storage.RepositoryInterface
	MaxPokemonId int
	Endpoint     string
	randFunc     RandFunc
}

const maxPokemonId = 1025

func NewPokemonService(config PokemonServiceConfig) *PokemonService {
	if config.randFunc == nil {
		config.randFunc = rand.Intn
	}

	if config.Endpoint == "" {
		config.Endpoint = "https://pokeapi.co/api/v2/pokemon/%d"
	}

	if config.MaxPokemonId == 0 || config.MaxPokemonId > maxPokemonId {
		config.MaxPokemonId = 1025
	}

	return &PokemonService{
		endpoint:     config.Endpoint,
		maxPokemonId: config.MaxPokemonId,
		repository:   config.Repository,
		randFunc:     config.randFunc,
	}
}

func (p *PokemonService) GetRandomPokemon() *Pokemon {
	id := p.randFunc(p.maxPokemonId)
	pokemon, _ := p.getPokemonById(id)

	p.savePokemon(id, pokemon)

	return pokemon
}

func (p *PokemonService) MarkCorrectGuess(pid int) {
	p.repository.Increment(pid, "correct")
}

func (p *PokemonService) MarkIncorrectGuess(pid int) {
	p.repository.Increment(pid, "incorrect")
}

func (p *PokemonService) getPokemonById(id int) (*Pokemon, error) {
	fromDb := p.getPokemonFromDb(id)

	if fromDb != nil {
		return fromDb, nil
	}

	fromApi, err := p.getPokemonFromApi(id)
	pokemon := &Pokemon{
		Id:    id,
		Name:  fromApi.Name,
		Image: fromApi.Sprites.FrontDefault,
	}
	return pokemon, err
}

func (p *PokemonService) getPokemonFromDb(id int) *Pokemon {
	if dto, err := p.repository.FindById(id); err == nil {
		p.repository.StoreHit(id)

		return &Pokemon{
			Id:    dto.Id,
			Name:  dto.Name,
			Image: dto.Image,
		}
	}

	return nil
}

func (p *PokemonService) savePokemon(id int, pokemon *Pokemon) {
	p.repository.Save(storage.PokemonDto{
		Id:    id,
		Name:  pokemon.Name,
		Image: pokemon.Image,
	})

	fmt.Println("Saved pokemon to db", pokemon)
}

func (p *PokemonService) getPokemonFromApi(id int) (PokemonRes, error) {
	resp, err := http.Get(fmt.Sprintf(p.endpoint, id))

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
