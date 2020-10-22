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
	CategoryModule struct {
		db   *sql.DB
		Name string
	}
)

func NewCategoryModule(db *sql.DB) *CategoryModule {
	return &CategoryModule{
		db:   db,
		Name: "module/Category",
	}
}

func (m CategoryModule) All(ctx context.Context, param model.AllCategory) ([]model.CategoryResponse, *Error) {
	var all []model.CategoryResponse

	data, err := (&model.Category{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all Category"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no Category found"
		}
		return []model.CategoryResponse{}, NewErrorWrap(err, m.Name, "all/Category",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m CategoryModule) Add(ctx context.Context, param model.Category) (model.CategoryResponse, *Error) {

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add Category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "add/Category",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m CategoryModule) One(ctx context.Context, param model.Category) (model.CategoryResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one Category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "one/Category",
			message, status)
	}

	return data.Response(), nil
}

func (m CategoryModule) Update(ctx context.Context, param model.Category) (model.CategoryResponse, *Error) {
	var emptyUUID uuid.UUID

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update Category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "update/Category",
			message, status)
	}
	return param.Response(), nil
}

func (m CategoryModule) Delete(ctx context.Context, param model.Category) (model.CategoryResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete Category"

		return model.CategoryResponse{}, NewErrorWrap(err, m.Name, "delete/Category",
			message, status)
	}
	return param.Response(), nil
}
