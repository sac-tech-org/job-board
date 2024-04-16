package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type AuthStore interface {
	GetProviders() []tpmodels.ProviderInput
	GetTypeInput() supertokens.TypeInput
	GetWebURI() string
}

type DataStore interface {
	CreateUser(context.Context, PostUserInput) error
	DeleteUser(context.Context, DeleteUserInput) error
	GetUser(context.Context, GetUserInput) (UserMetadata, error)
	GetUserList(context.Context, GetUserListInput) ([]UserMetadata, error)
	PutUser(ctx context.Context, first, id, last, username string) error
}

type IdentityStore interface {
	GetUser(string) (UserIdentity, error)
}

type Server struct {
	authStore AuthStore
	dataStore DataStore
	idStore   IdentityStore
	router    http.Handler
}

type HTTPError struct {
	status  int
	message string
}

func (e HTTPError) Error() string { return e.message }

func respond(w http.ResponseWriter, r *http.Request, data any, err error, status int) {
	type response struct {
		Data  any    `json:"data,omitempty"`
		Error string `json:"error,omitempty"`
	}

	w.Header().Set("Content-Type", "application/json")
	out := response{}

	if err != nil {
		var hErr HTTPError
		if errors.As(err, &hErr) {
			w.WriteHeader(hErr.status)
			out.Error = hErr.message
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			out.Error = "unknown error"
		}
	}

	w.WriteHeader(status)
	out.Data = data

	_ = json.NewEncoder(w).Encode(out)
}

func (s *Server) Router() http.Handler {
	return s.router
}

func (s *Server) routes() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{s.authStore.GetWebURI()},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   append([]string{"Content-Type"}, supertokens.GetAllCORSHeaders()...),
		AllowCredentials: true,
	}))
	r.Use(supertokens.Middleware)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
	})

	r.Route("/user", func(r chi.Router) {
		r.Get("/", s.handleGetUserList)
		r.Post("/", s.handlePostUser)

		r.With(withUserSession(true)).Get("/me", s.handleGetCurrentUser)

		r.Route("/{username}", func(r chi.Router) {
			r.Use(withUserSession(false))
			r.Get("/", s.handleGetUser)
			r.Group(func(r chi.Router) {
				r.Use(s.withAccountOwner("username"))
				r.Delete("/", s.handleDeleteUser)
				r.Put("/", s.handlePutUser)
			})
		})
	})

	s.router = r
}

func NewServer(a AuthStore, d DataStore, i IdentityStore) (Server, error) {
	s := Server{
		authStore: a,
		dataStore: d,
		idStore:   i,
	}

	if err := s.superTokenInit(); err != nil {
		return Server{}, fmt.Errorf("error initializing SuperTokens: %w", err)
	}

	s.routes()

	return s, nil
}

func (s *Server) superTokenInit() error {
	// https://supertokens.com/docs/thirdpartyemailpassword/pre-built-ui/setup/backend
	return supertokens.Init(s.authStore.GetTypeInput())
}
