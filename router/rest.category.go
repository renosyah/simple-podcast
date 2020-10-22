package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/simple-podcast/api"
	"github.com/renosyah/simple-podcast/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddCategory(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Category

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return categoryModule.Add(ctx, param)
}

func HandlerAllCategory(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.AllCategory

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return categoryModule.All(ctx, param)
}

func HandlerOneCategory(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return categoryModule.One(ctx, model.Category{ID: id})
}

func HandlerUpdateCategory(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Category

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return categoryModule.Update(ctx, param)
}

func HandlerDeleteCategory(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Category/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return categoryModule.Delete(ctx, model.Category{ID: id})
}
