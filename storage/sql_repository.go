package storage

import (
	"database/sql"
	"fmt"
	"sort"

	_ "github.com/glebarez/go-sqlite"
)

const file string = "database.db"

type SqlRepository struct {
	db *sql.DB
}

func NewSqlRepository() *SqlRepository {
	db, err := sql.Open("sqlite", file)

	if err != nil {
		panic(err)
	}

	if _, err := db.Exec("SELECT * FROM pokemon"); err != nil {
		panic(err)
	}

	return &SqlRepository{
		db,
	}
}

func (s *SqlRepository) FindById(id int) (PokemonDto, error) {
	fmt.Println("searching database for id: ", id)

	row := s.db.QueryRow(`
		SELECT id, name, image, incorrect, correct
		FROM pokemon
		WHERE id = ?
	`, id)

	dto := PokemonDto{}
	var err error

	if err = row.Scan(&dto.Id, &dto.Name, &dto.Image, &dto.Incorrect, &dto.Correct); err == sql.ErrNoRows {
		return PokemonDto{}, err
	} else if err != nil {
		panic(err)
	}

	return dto, nil
}

func (s *SqlRepository) Increment(id int, column string) error {
	fmt.Println("incrementing column: ", column, " for id: ", id)
	query := fmt.Sprintf("UPDATE pokemon SET %s = %s + 1 WHERE id = ?", column, column)
	_, err := s.db.Exec(query, id)
	if err != nil {
		panic(err)
	}
	return err
}

func (s *SqlRepository) StoreHit(id int) error {
	return s.Increment(id, "hits")
}

func (s *SqlRepository) Save(p PokemonDto) error {
	_, err := s.db.Exec(`
		INSERT INTO pokemon (id, name, image, incorrect, correct, hits)
		VALUES (?, ?, ?, 0, 0, 1)
	`, p.Id, p.Name, p.Image)

	if err != nil {
		return nil
	}

	return nil
}

func (s *SqlRepository) All() []PokemonDto {
	rows, err := s.db.Query(`
		SELECT id, name, image, incorrect, correct
		FROM pokemon
	`)

	if err != nil {
		panic(err)
	}

	pokemon := []PokemonDto{}

	defer rows.Close()

	for rows.Next() {
		dto := PokemonDto{}
		_ = rows.Scan(&dto.Id, &dto.Name, &dto.Image, &dto.Incorrect, &dto.Correct)
		pokemon = append(pokemon, dto)
	}

	sort.Slice(pokemon, func(i, j int) bool {
		return pokemon[i].Correct > pokemon[j].Correct
	})

	return pokemon
}
