package users

import (
	"database/sql"
	"fmt"

	"log"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	repo := &Repository{}
	return repo
}

func (r *Repository) Connect(db *sql.DB) error {
	r.db = db
	return r.migrate()
}

func (r *Repository) migrate() error {
	_, err := r.db.Exec(`CREATE TABLE IF NOT EXISTS resource (

	)`)
	if err != nil {
		log.Printf("Failed to create users table: %v", err)
		return fmt.Errorf("failed to create users table: %w", err)
	}
	fmt.Println("Users table migrated successfully")
	return nil
}
