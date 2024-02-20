package pokemon

import (
	"testing"

	"github.com/jaymo107/whos-that-pokemon/storage"
)

func TestGetsExistingPokemonFromTheDatabase(t *testing.T) {
	config := PokemonServiceConfig{
		randFunc: func(int) int {
			return 1
		},
		Repository: &storage.FakeRepository{},
	}

	// Mock get handler
}

// func TestGetsPokemonFromTheApi(t *testing.T) {
// 	// Mock out the repository
// }

// func TestMarksGuessAsCorrect(t *testing.T) {
// 	// Mock out the repository
// }

// func TestMarksGuessAsIncorrect(t *testing.T) {
// 	// Mock out the repository
// }
