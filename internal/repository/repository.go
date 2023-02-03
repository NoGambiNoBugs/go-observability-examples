package repository

import (
	"context"
	"database/sql"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
)

// Repository is the main repo of app.
type Repository struct {
	db *sql.DB
}

// InsertCustomer inserts a new customer.
func (r Repository) InsertCustomer(ctx context.Context, customer entity.Customer) (err error) {
	_, err = r.db.ExecContext(ctx, "INSERT INTO customer VALUES ($1, $2, $3)",
		customer.ID,
		customer.Name,
		customer.Email,
	)
	return
}

// New returns a instance of Repository.
func New(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}
