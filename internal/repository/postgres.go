package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

func NewPostgresRepo(conn string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	r := &PostgresRepo{db: db}
	r.init()

	return r, nil
}

func (r *PostgresRepo) init() {
	r.db.Exec(`
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT
	)
		`)
}

func (r *PostgresRepo) CreateItem(name string) Item {
	var id int
	r.db.QueryRow(
		"INSERT INTO items (name) VALUES ($1) RETURNING id",
		name,
	).Scan(&id)
	return Item{ID: id, Name: name}
}

func (r *PostgresRepo) ListItems() []Item {
	rows, _ := r.db.Query("SELECT id, name FROM items")

	var items []Item

	for rows.Next() {
		var i Item
		rows.Scan(&i.ID, &i.Name)
		items = append(items, i)
	}

	return items
}
