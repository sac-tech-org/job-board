package server

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type (
	DeleteUserInput struct {
		UUID uuid.UUID `json:"id"`
	}

	GetUserInput struct {
		UUID uuid.UUID `json:"id"`
	}

	GetUserListInput struct{}

	PostUserInput struct {
		Email     string `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Username  string `json:"username"`
	}

	PutUserInput struct {
		Email string `json:"email,omitempty"`
		UUID  uuid.UUID
		Name  string `json:"name,omitempty"`
	}

	User struct {
		Email     string    `json:"email"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		UUID      uuid.UUID `json:"uuid"`
		Username  string    `json:"username"`
	}
)

var UserNotFound = HTTPError{http.StatusNotFound, "User not found"}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[uuid.UUID](ctx, userIDCTXKey)

	if err := s.dataStore.DeleteUser(ctx, DeleteUserInput{id}); err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, nil, nil, http.StatusNoContent)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[uuid.UUID](ctx, userIDCTXKey)

	out, err := s.dataStore.GetUser(ctx, GetUserInput{id})
	if err != nil {
		if errors.Is(err, UserNotFound) {
			respond(w, r, nil, UserNotFound, 0)
			return
		}

		respond(w, r, nil, HTTPError{status: http.StatusInternalServerError, message: err.Error()}, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handleGetUserList(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	out, err := s.dataStore.GetUserList(ctx, GetUserListInput{})
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handlePostUser(w http.ResponseWriter, r *http.Request) {
	var in PostUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	ctx := r.Context()

	out, err := s.dataStore.PostUser(ctx, in)
	if err != nil {
		log.Printf("err: %v\n", err)
		respond(w, r, nil, HTTPError{status: http.StatusInternalServerError, message: err.Error()}, 0)
		return
	}

	respond(w, r, out, nil, http.StatusCreated)
}

func (s *Server) handlePutUser(w http.ResponseWriter, r *http.Request) {
	var in PutUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	ctx := r.Context()
	in.UUID = *ContextValue[uuid.UUID](ctx, userIDCTXKey)

	out, err := s.dataStore.PutUser(ctx, in)
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, out, nil, http.StatusOK)
}
