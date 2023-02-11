package port

import (
	"context"

	"github.com/NoGambiNoBugs/go-observability-examples/internal/entity"
)

//go:generate gowrap gen -g -i CustomerUsecase -t ./templates/log_template.go.tmpl -o ./decorators/log/customer_usecase_with_log.go
//go:generate gowrap gen -g -i CustomerUsecase -t ./templates/red_template.go.tmpl -o ./decorators/red/customer_usecase_with_red.go -v "Namespace=example"
// CustomerUsecase is the contract for the customer usecase.
type CustomerUsecase interface {
	Create(ctx context.Context, customer entity.Customer) (err error)
}
