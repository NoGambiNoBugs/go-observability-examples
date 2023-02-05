package bootstrap

import (
	"net/http"
	"os"

	_ "github.com/NoGambiNoBugs/go-observability-examples/internal/env" //no lint
	"github.com/NoGambiNoBugs/go-observability-examples/internal/handler"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/instrumentation"
	logDecorator "github.com/NoGambiNoBugs/go-observability-examples/internal/port/decorators/log"
	redDecorator "github.com/NoGambiNoBugs/go-observability-examples/internal/port/decorators/red"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/repository"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/tools/postgres"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/usecase"
)

type App struct {
	server *http.Server
}

type customerEndpoint struct {
	h handler.Handler
}

func (c *customerEndpoint) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch {
	case r.Method == http.MethodPost:
		c.h.PostCustomer(w, r)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("not found"))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func route(h handler.Handler) *http.Server {
	mux := http.NewServeMux()
	mux.Handle("/customers/", &customerEndpoint{h})
	mux.Handle("/metrics", instrumentation.RegistryHandler())

	return &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: mux,
	}
}

func (s App) Run() error {
	return s.server.ListenAndServe()
}

func Setup() (App, error) {
	instrumentation.Init()

	db, err := postgres.Init()
	if err != nil {
		return App{}, err
	}

	subsystem := "api"

	repoWithRED, err := redDecorator.NewRepositoryWithREDHistogram(repository.New(db), subsystem, nil)
	if err != nil {
		return App{}, err
	}
	repo := logDecorator.NewRepositoryWithLog(repoWithRED)

	usecaseWithRED, err := redDecorator.NewCustomerUsecaseWithREDHistogram(usecase.New(repo), subsystem, nil)
	if err != nil {
		return App{}, err
	}
	usecase := logDecorator.NewCustomerUsecaseWithLog(usecaseWithRED)

	h := handler.New(usecase)

	return App{
		server: route(h),
	}, nil
}
