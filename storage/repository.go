package storage

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PokemonDto struct {
	Id        int
	Name      string
	Image     string
	Incorrect int
	Correct   int
}

func (p *PokemonDto) FormatName() *PokemonDto {
	p.Name = cases.Title(language.English).String(p.Name)
	return p
}

type RepositoryInterface interface {
	FindById(id int) (PokemonDto, error)
	Save(pokemon PokemonDto) error
	Increment(id int, column string) error
	StoreHit(id int) error
	All() []PokemonDto
}
