package server_test

import (
	"context"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/rusher2004/job-board/api/server"
	"github.com/stretchr/testify/assert"
)

type mockUserGetter struct {
	server.DataStore

	user server.User
	err  error
}

func (m mockUserGetter) GetUser(context.Context, server.GetUserInput) (server.User, error) {
	return m.user, m.err
}

func TestGetUser(t *testing.T) {
	a := assert.New(t)

	var userUUID = uuid.MustParse("07a86fdf-fe39-45ea-af57-e4aa9e8068b0")

	tests := []struct {
		exp        string
		name       string
		path       string
		statusCode int
		store      server.DataStore
	}{
		// happy path
		{
			name:       "User is returned",
			exp:        `{"data":{"email":"robert@test.com","id":"07a86fdf-fe39-45ea-af57-e4aa9e8068b0","name":"Robert"}}`,
			path:       "/user/" + userUUID.String(),
			statusCode: 200,
			store: mockUserGetter{
				user: server.User{
					Email:     "robert@test.com",
					UUID:      userUUID,
					FirstName: "Robert",
				},
			},
		},
		// not found
		{
			name:       "User not found",
			exp:        `{"error":"User not found"}`,
			path:       "/user/" + userUUID.String(),
			statusCode: 404,
			store: mockUserGetter{
				err: server.UserNotFound,
			},
		},
	}

	for _, tt := range tests {
		r := httptest.NewRequest("GET", tt.path, nil)
		w := httptest.NewRecorder()

		s := server.NewServer(tt.store)
		s.Router().ServeHTTP(w, r)

		res := w.Result()
		b, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %v", err)
		}

		a.JSONEqf(
			tt.exp,
			string(b),
			"%s-%s:\n got: %s\n want: %s\n", t.Name(), tt.name, string(b), tt.exp,
		)
		a.Equalf(
			tt.statusCode,
			res.StatusCode,
			"%s-%s:\n got status code: %d\n want status code: %d\n", t.Name(), tt.name, res.StatusCode, tt.statusCode,
		)
	}
}
