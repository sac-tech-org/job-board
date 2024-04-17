package server

import (
	"encoding/json"
	"net/http"
)

type (
	OrganizationResource struct {
		Description string `json:"description"`
		ID          string `json:"id"`
		Name        string `json:"name"`
		Website     string `json:"website"`
	}

	PostOrganizationInput struct {
		Description string `json:"description"`
		ID          string
		Name        string `json:"name"`
		Website     string `json:"website"`
	}
)

func (s *Server) handlePostOrganization(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := ContextValue[string](ctx, userIDCTXKey)

	var in PostOrganizationInput
	if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
		hErr := HTTPError{http.StatusUnprocessableEntity, "invalid JSON"}
		respond(w, r, nil, hErr, 0)
		return
	}
	in.ID = id

	// TODO: bring in a validator. All other handlers should have one too.

	oid, err := s.dataStore.CreateOrganization(ctx, in.Description, in.Name, id, in.Website)
	if err != nil {
		hErr := HTTPError{http.StatusInternalServerError, "error creating organization: " + err.Error()}
		respond(w, r, nil, hErr, 0)
		return
	}

	out := OrganizationResource{
		Description: in.Description,
		ID:          oid,
		Name:        in.Name,
		Website:     in.Website,
	}

	respond(w, r, out, nil, http.StatusCreated)
}
