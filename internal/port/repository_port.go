package port

import (
	"context"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
)

//go:generate gowrap gen -g -i Repository -t ./templates/log_template.go.tmpl -o ./decorators/log/repository_with_log.go
// Repository is the contract for the repository.
type Repository interface {
	InsertCustomer(ctx context.Context, customer entity.Customer) (err error)
}
