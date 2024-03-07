package server

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// contextKey for key/vals held in request context
type contextKey struct {
	name string
}

func (k *contextKey) String() string {
	return "context key value: " + k.name
}

var userIDCTXKey = &contextKey{"userID"}

// ContextValue is a shortcut to fetch the value of type T from context.
func ContextValue[T any](ctx context.Context, key *contextKey) *T {
	val, _ := ctx.Value(key).(*T)
	return val
}

func withUserID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		u, err := uuid.Parse(id)
		if err != nil {
			hErr := HTTPError{http.StatusUnprocessableEntity, "invalid id format"}
			respond(w, r, nil, hErr, hErr.status)
			return
		}

		ctx := context.WithValue(r.Context(), userIDCTXKey, &u)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
