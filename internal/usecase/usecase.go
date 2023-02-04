package usecase

import (
	"context"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/port"
	"github.com/google/uuid"
)

// CustomerUsecase manage the usecases for customer
type CustomerUsecase struct {
	repo port.Repository
}

// Create a new customer.
func (c CustomerUsecase) Create(ctx context.Context, customer entity.Customer) error {
	customer.ID = uuid.New()
	return c.repo.InsertCustomer(ctx, customer)
}

// New returns a instance of CustomerUsecase.
func New(repo port.Repository) CustomerUsecase {
	return CustomerUsecase{
		repo: repo,
	}
}
