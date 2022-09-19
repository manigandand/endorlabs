package api

import (
	"net/http"

	"github.com/manigandand/adk/api"
	"github.com/manigandand/adk/errors"
	"github.com/manigandand/adk/respond"
	"github.com/manigandand/endorlabs/schema"
)

func savePersonHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()

	var req schema.Person
	if err := api.Decode(r, &req); err != nil {
		return err
	}

	if err := store.Store(ctx, &req); err != nil {
		return errors.InternalServer(err.Error())
	}

	return respond.Created(w, &req)
}

func saveAnimalsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()

	var req schema.Animal
	if err := api.Decode(r, &req); err != nil {
		return err
	}

	if err := store.Store(ctx, &req); err != nil {
		return errors.InternalServer(err.Error())
	}

	return respond.Created(w, &req)
}
