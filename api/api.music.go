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
	MusicModule struct {
		db   *sql.DB
		Name string
	}
)

func NewMusicModule(db *sql.DB) *MusicModule {
	return &MusicModule{
		db:   db,
		Name: "module/Music",
	}
}

func (m MusicModule) All(ctx context.Context, param model.AllMusic) ([]model.MusicResponse, *Error) {
	var all []model.MusicResponse

	data, err := (&model.Music{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Music"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Music found"
		}
		return []model.MusicResponse{}, NewErrorWrap(err, m.Name, "all/Music",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m MusicModule) Add(ctx context.Context, param model.Music) (model.MusicResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Music"

		return model.MusicResponse{}, NewErrorWrap(err, m.Name, "add/Music",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m MusicModule) One(ctx context.Context, param model.Music) (model.MusicResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Music"

		return model.MusicResponse{}, NewErrorWrap(err, m.Name, "one/Music",
			message, status)
	}

	return data.Response(), nil
}

func (m MusicModule) Update(ctx context.Context, param model.Music) (model.MusicResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Music"

		return model.MusicResponse{}, NewErrorWrap(err, m.Name, "update/Music",
			message, status)
	}
	return param.Response(), nil
}

func (m MusicModule) Delete(ctx context.Context, param model.Music) (model.MusicResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Music"

		return model.MusicResponse{}, NewErrorWrap(err, m.Name, "delete/Music",
			message, status)
	}
	return param.Response(), nil
}
