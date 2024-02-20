package storage

import (
	_ "github.com/glebarez/go-sqlite"
)

type FakeRepository struct {
	//
}

func (f *FakeRepository) FindById(id int) (PokemonDto, error) {
	return PokemonDto{}, nil
}

func (f *FakeRepository) Increment(id int, column string) error {
	return nil
}

func (f *FakeRepository) StoreHit(id int) error {
	return nil
}

func (f *FakeRepository) Save(p PokemonDto) error {
	return nil
}

func (f *FakeRepository) All() []PokemonDto {
	return []PokemonDto{}
}
