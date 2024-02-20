package storage

import "testing"

func TestFormatsName(t *testing.T) {
	names := map[string]string{
		"bulbasaur": "Bulbasaur",
		"ivysaur":   "Ivysaur",
		"venusaur":  "Venusaur",
	}

	for name, expexted := range names {
		p := PokemonDto{Name: name}
		p.FormatName()
		if p.Name != expexted {
			t.Errorf("expected %s, got %s", expexted, p.Name)
		}
	}
}
