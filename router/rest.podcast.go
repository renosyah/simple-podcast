package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/renosyah/simple-podcast/api"
	"github.com/renosyah/simple-podcast/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerAddPodcast(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()

	var param model.Podcast

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/create/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return podcastModule.Add(ctx, param)
}

func HandlerAllPodcast(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	var param model.AllPodcast

	err := ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/all/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	return podcastModule.All(ctx, param)
}

func HandlerOnePodcast(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/detail"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return podcastModule.One(ctx, model.Podcast{ID: id})
}

func HandlerUpdatePodcast(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	var param model.Podcast

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/update"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	err = ParseBodyData(ctx, r, &param)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/update/param"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}

	param.ID = id

	return podcastModule.Update(ctx, param)
}

func HandlerDeletePodcast(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Podcast/delete"),
			http.StatusText(http.StatusNotFound), http.StatusNotFound)
	}

	return podcastModule.Delete(ctx, model.Podcast{ID: id})
}
