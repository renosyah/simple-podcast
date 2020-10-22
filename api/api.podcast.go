package api

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/pkg/errors"
	"github.com/renosyah/simple-podcast/model"
	uuid "github.com/satori/go.uuid"
)

type (
	PodcastModule struct {
		db   *sql.DB
		Name string
	}
)

func NewPodcastModule(db *sql.DB) *PodcastModule {
	return &PodcastModule{
		db:   db,
		Name: "module/Podcast",
	}
}

func (m PodcastModule) All(ctx context.Context, param model.AllPodcast) ([]model.PodcastResponse, *Error) {
	var all []model.PodcastResponse

	data, err := (&model.Podcast{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Podcast"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Podcast found"
		}
		return []model.PodcastResponse{}, NewErrorWrap(err, m.Name, "all/Podcast",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m PodcastModule) Add(ctx context.Context, param model.Podcast) (model.PodcastResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Podcast"

		return model.PodcastResponse{}, NewErrorWrap(err, m.Name, "add/Podcast",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m PodcastModule) One(ctx context.Context, param model.Podcast) (model.PodcastResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Podcast"

		return model.PodcastResponse{}, NewErrorWrap(err, m.Name, "one/Podcast",
			message, status)
	}

	return data.Response(), nil
}

func (m PodcastModule) Update(ctx context.Context, param model.Podcast) (model.PodcastResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Podcast"

		return model.PodcastResponse{}, NewErrorWrap(err, m.Name, "update/Podcast",
			message, status)
	}
	return param.Response(), nil
}

func (m PodcastModule) Delete(ctx context.Context, param model.Podcast) (model.PodcastResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Podcast"

		return model.PodcastResponse{}, NewErrorWrap(err, m.Name, "delete/Podcast",
			message, status)
	}
	return param.Response(), nil
}
