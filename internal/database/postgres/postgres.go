package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sureshchandak1/go-jwt/internal/config"
)

type Postgres struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Postgres, error) {

	p := cfg.PostgresConfig
	connectionUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		p.Host, p.Port, p.Username, p.Password, p.Dbname, p.Schema)

	db, err := sql.Open("postgres", connectionUrl)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT,
		age INTEGER
		)`)

	if err != nil {
		return nil, err
	}

	return &Postgres{
		Db: db,
	}, nil
}
