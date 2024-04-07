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
	DeleteUser(context.Context, DeleteUserInput) error
	GetUser(context.Context, GetUserInput) (User, error)
	GetUserList(context.Context, GetUserListInput) ([]User, error)
	PostUser(context.Context, PostUserInput) (User, error)
	PutUser(context.Context, PutUserInput) (User, error)
}

type UserStore interface {
	GetMetadata(string) (UserMetadata, error)
	UpdateMetadata(string, UserMetadata) (UserMetadata, error)
}

type Server struct {
	authStore AuthStore
	dataStore DataStore
	router    http.Handler
	userStore UserStore
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

	r.Route("/user", func(userRoute chi.Router) {
		userRoute.Get("/", s.handleGetUserList)
		userRoute.Post("/", s.handlePostUser)

		userRoute.Route("/{id}", func(userIDRoute chi.Router) {
			// userIDRoute.Use(withUserSession)
			userIDRoute.Use(withUserID)
			userIDRoute.Delete("/", s.handleDeleteUser)
			userIDRoute.Get("/", s.handleGetUser)
			userIDRoute.Put("/", s.handlePutUser)
		})
	})

	s.router = r
}

func NewServer(a AuthStore, d DataStore, u UserStore) (Server, error) {
	s := Server{
		authStore: a,
		dataStore: d,
		userStore: u,
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
