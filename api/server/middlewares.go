package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/session/sessmodels"
	"github.com/supertokens/supertokens-golang/supertokens"
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

// withUserSession will check for a session and add the userID to the request context. If no session
// exists the request will continue without a userID in the context.
func withUserSession(required bool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sess, err := session.GetSession(r, w, &sessmodels.VerifySessionOptions{
				SessionRequired: &required,
			})
			if err != nil {
				// let supertokens handle its errors, including session not found when required
				if err := supertokens.ErrorHandler(err, r, w); err != nil {
					// catch some other unknown errors
					hErr := HTTPError{http.StatusInternalServerError, "error getting session"}
					respond(w, r, nil, hErr, hErr.status)
				}
				return
			}

			if sess != nil {
				// session exists, add userID to context
				id := sess.GetUserID()
				ctx := context.WithValue(r.Context(), userIDCTXKey, &id)

				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			// no session, continue
			next.ServeHTTP(w, r)
		})
	}
}

// withAccountOwner will validate that the user id in the request context is the same account as the
// named resource. It must be used after withUserSession.
func (s *Server) withAccountOwner(param string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			id := *ContextValue[string](ctx, userIDCTXKey)

			p := chi.URLParam(r, param)
			if p == "" {
				hErr := HTTPError{http.StatusBadRequest, fmt.Sprintf("missing {%s} parameter", param)}
				respond(w, r, nil, hErr, hErr.status)
				return
			}

			u, err := s.dataStore.GetUser(ctx, GetUserInput{ID: id})
			if err != nil {
				respond(w, r, nil, err, 0)
				return
			}

			b, err := json.Marshal(u)
			if err != nil {
				hErr := HTTPError{http.StatusInternalServerError, "error marshalling accountOwner data"}
				respond(w, r, nil, hErr, hErr.status)
				return
			}

			var m map[string]any
			if err := json.Unmarshal(b, &m); err != nil {
				hErr := HTTPError{http.StatusInternalServerError, "error unmarshalling accountOwner data"}
				respond(w, r, nil, hErr, hErr.status)
				return
			}

			if m[param] != p {
				hErr := HTTPError{http.StatusForbidden, "accountOwner does not match request"}
				respond(w, r, nil, hErr, hErr.status)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
