package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
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

func getObjectByIDHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	objectID := chi.URLParam(r, "objectID")
	if objectID == "" {
		return errors.BadRequest("object id required")
	}

	obj, err := store.GetObjectByID(ctx, objectID)
	if err != nil {
		return errors.InternalServer(err.Error())
	}

	respond.OK(w, obj)
	return nil
}

func getObjectByNameHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	objectName := chi.URLParam(r, "objectName")
	if objectName == "" {
		return errors.BadRequest("object name required")
	}

	obj, err := store.GetObjectByName(ctx, objectName)
	if err != nil {
		return errors.InternalServer(err.Error())
	}

	respond.OK(w, obj)
	return nil
}

func listObjectsHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	kindName := chi.URLParam(r, "kindName")
	if kindName == "" {
		return errors.BadRequest("object id required")
	}

	obj, err := store.ListObjects(ctx, kindName)
	if err != nil {
		return errors.InternalServer(err.Error())
	}

	respond.OK(w, obj)
	return nil
}

func deleteObjectHandler(w http.ResponseWriter, r *http.Request) *errors.AppError {
	ctx := r.Context()
	objectID := chi.URLParam(r, "objectID")
	if objectID == "" {
		return errors.BadRequest("object id required")
	}

	if err := store.DeleteObject(ctx, objectID); err != nil {
		return errors.InternalServer(err.Error())
	}

	respond.NoContent(w, nil)
	return nil
}
