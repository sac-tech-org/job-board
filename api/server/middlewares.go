package server

import (
	"context"
	"net/http"

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
