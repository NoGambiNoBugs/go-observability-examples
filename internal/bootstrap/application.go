package bootstrap

import (
	"net/http"
	"os"

	_ "github.com/NoGambiNoBugs/go-observability-examples/internal/env" //no lint
	"github.com/NoGambiNoBugs/go-observability-examples/internal/handler"
	logDecorator "github.com/NoGambiNoBugs/go-observability-examples/internal/port/decorators/log"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/repository"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/tools/postgres"
	"github.com/NoGambiNoBugs/go-observability-examples/internal/usecase"
	"github.com/rs/zerolog"
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

	return &http.Server{
		Addr:    os.Getenv("SERVER_ADDRESS"),
		Handler: mux,
	}
}

func (s App) Run() error {
	return s.server.ListenAndServe()
}

func Setup() (App, error) {
	db, err := postgres.Init()
	if err != nil {
		return App{}, err
	}

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Logger()
	zerolog.DefaultContextLogger = &logger

	repo := logDecorator.NewRepositoryWithLog(repository.New(db))
	usecase := logDecorator.NewCustomerUsecaseWithLog(usecase.New(repo))
	h := handler.New(usecase)

	return App{
		server: route(h),
	}, nil
}
