package repository

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sql.DB
}

type Item struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPostgresRepo(conn string) (*PostgresRepo, error) {
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	repo := &PostgresRepo{db: db}
	if err := repo.InitSchema(); err != nil {
		return nil, err
	}

	return repo, nil
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
	var item Item

	query := `
	INSERT INTO items (name)
	VALUES ($1)
	RETURNING id, name, created_at
	`

	r.db.QueryRow(query, name).Scan(
		&item.ID,
		&item.Name,
		&item.CreatedAt,
	)

	return item
}

func (r *PostgresRepo) ListItems() []Item {
	rows, _ := r.db.Query(
		"SELECT id, name, created_at FROM items ORDER BY id",
	)
	defer rows.Close()

	var items []Item

	for rows.Next() {
		var item Item
		rows.Scan(&item.ID, &item.Name, &item.CreatedAt)
		items = append(items, item)
	}

	return items
}

func (r *PostgresRepo) InitSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL
	)
	`
	_, err := r.db.Exec(query)
	return err
}

func (r *PostgresRepo) AddItem(name string) error {
	_, err := r.db.Exec("INSERT INTO items (name) VALUES ($1)", name)
	return err
}

func (r *PostgresRepo) GetItems() ([]string, error) {
	rows, err := r.db.Query("SELECT name FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}

	return items, nil
}

func (r *PostgresRepo) Init() error {
	query := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err := r.db.Exec(query)
	return err
}
