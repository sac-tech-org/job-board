package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/sync/errgroup"
)

type (
	DeleteUserInput struct {
		ID string `json:"id"`
	}

	Email struct {
		Address  string `json:"address"`
		Verified bool   `json:"verified"`
	}

	GetUserInput struct {
		ID string `json:"id"`
	}

	UserResource struct {
		Email     Email  `json:"email"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Username  string `json:"username"`
	}

	GetUserListInput struct{}

	PostUserInput struct {
		FirstName string `json:"firstName"`
		ID        string `json:"id"`
		LastName  string `json:"lastName"`
		Username  string `json:"username"`
	}

	PutUserInput struct {
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		Username  string `json:"username,omitempty"`
	}

	UserIdentity struct {
		Email Email
		ID    string
	}

	UserMetadata struct {
		FirstName string `json:"firstName,omitempty"`
		LastName  string `json:"lastName,omitempty"`
		Username  string `json:"username,omitempty"`
	}
)

var UserNotFound = HTTPError{http.StatusNotFound, "User not found"}

func (s *Server) handleDeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[string](ctx, userIDCTXKey)

	if err := s.dataStore.DeleteUser(ctx, DeleteUserInput{id}); err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	respond(w, r, nil, nil, http.StatusNoContent)
}

func (s *Server) handleGetCurrentUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[string](ctx, userIDCTXKey)

	ui, err := s.idStore.GetUser(id)
	if err != nil {
		respond(w, r, nil, err, 0)
		return
	}

	um, err := s.dataStore.GetUser(ctx, GetUserInput{ID: ui.ID})
	if err != nil {
		hErr := HTTPError{http.StatusInternalServerError, "error getting user: " + err.Error()}
		respond(w, r, nil, hErr, 0)
		return
	}

	out := UserResource{
		Email:     ui.Email,
		FirstName: um.FirstName,
		LastName:  um.LastName,
		Username:  um.Username,
	}

	respond(w, r, out, nil, http.StatusOK)
}

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request) {}

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
	ctx := r.Context()

	var in PostUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	if err := s.dataStore.CreateUser(ctx, in); err != nil {
		hErr := HTTPError{http.StatusInternalServerError, "error creating user: " + err.Error()}
		respond(w, r, nil, hErr, 0)
		return
	}

	out := UserResource{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
	}

	respond(w, r, out, nil, http.StatusCreated)
}

func (s *Server) handlePutUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := *ContextValue[string](ctx, userIDCTXKey)

	var in PutUserInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}

	g, errCTX := errgroup.WithContext(ctx)
	g.SetLimit(2)

	var ui UserIdentity
	g.Go(func() error {
		var err error
		if ui, err = s.idStore.GetUser(id); err != nil {
			return fmt.Errorf("error getting user identity: %w", err)
		}

		return nil
	})

	g.Go(func() error {
		if err := s.dataStore.PutUser(errCTX, in.FirstName, id, in.LastName, in.Username); err != nil {
			return fmt.Errorf("error updating user in db: %w", err)
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		hErr := HTTPError{http.StatusInternalServerError, "error updating user: " + err.Error()}
		respond(w, r, nil, hErr, 0)
		return
	}

	out := UserResource{
		Email:     ui.Email,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Username:  in.Username,
	}

	respond(w, r, out, nil, http.StatusOK)
}
