package port

import (
	"context"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
)

//go:generate gowrap gen -g -i Repository -t ./templates/log_template.go.tmpl -o ./decorators/log/repository_with_log.go
//go:generate gowrap gen -g -i Repository -t ./templates/red_histogram_template.go.tmpl -o ./decorators/red/repository_with_red_histogram.go -v "Namespace=example"
// Repository is the contract for the repository.
type Repository interface {
	InsertCustomer(ctx context.Context, customer entity.Customer) (err error)
}
