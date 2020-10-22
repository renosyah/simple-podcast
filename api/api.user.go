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
	UserModule struct {
		db   *sql.DB
		Name string
	}
)

func NewUserModule(db *sql.DB) *UserModule {
	return &UserModule{
		db:   db,
		Name: "module/User",
	}
}

func (m UserModule) All(ctx context.Context, param model.AllUser) ([]model.UserResponse, *Error) {
	var all []model.UserResponse

	data, err := (&model.User{}).All(ctx, m.db, param)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on query all User"

		if errors.Cause(err) == sql.ErrNoRows {
			status = http.StatusOK
			message = "no User found"
		}
		return []model.UserResponse{}, NewErrorWrap(err, m.Name, "all/User",
			message, status)
	}
	for _, each := range data {
		all = append(all, each.Response())
	}
	return all, nil

}
func (m UserModule) Add(ctx context.Context, param model.User) (model.UserResponse, *Error) {

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash User password"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "add/User",
			message, status)
	}

	param.Password = hash

	id, err := param.Add(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on add User"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "add/User",
			message, status)
	}
	param.ID = id
	return param.Response(), nil
}

func (m UserModule) One(ctx context.Context, param model.User) (model.UserResponse, *Error) {
	data, err := param.One(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on get one User"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "one/User",
			message, status)
	}

	return data.Response(), nil
}

func (m UserModule) Login(ctx context.Context, param model.User) (model.UserAuth, *Error) {
	data, err := param.OneByPhoneNumber(ctx, m.db)
	if err != nil {
		status := http.StatusInternalServerError
		message := "phone number or password invalid"

		return model.UserAuth{}, NewErrorWrap(err, m.Name, "login/User",
			message, status)
	}

	err = ComparePassword(ctx, data.Password, param.Password)
	if err != nil {
		status := http.StatusUnauthorized
		message := "phone number or password invalid"

		return model.UserAuth{}, NewErrorWrap(errors.New(message), m.Name, "login/User",
			message, status)
	}

	resp := model.UserAuth{
		SessionID:    uuid.NewV4(),
		UserResponse: data.Response(),
	}

	return resp, nil
}

func (m UserModule) Register(ctx context.Context, param model.User) (model.UserAuth, *Error) {
	var emptyUUID uuid.UUID

	cs, errCs := param.OneByPhoneNumber(ctx, m.db)
	if errCs != nil && errors.Cause(errCs) != sql.ErrNoRows {
		status := http.StatusInternalServerError
		message := "error on get one User"

		return model.UserAuth{}, NewErrorWrap(errCs, m.Name, "register/User",
			message, status)
	}

	if cs.ID != emptyUUID {
		status := http.StatusInternalServerError
		message := "User with this phone number exist"

		return model.UserAuth{}, NewErrorWrap(errors.New(message), m.Name, "register/User",
			message, status)
	}

	data, err := m.Add(ctx, param)
	if err != nil {
		return model.UserAuth{}, err
	}

	resp := model.UserAuth{
		SessionID:    uuid.NewV4(),
		UserResponse: data,
	}

	return resp, nil
}

func (m UserModule) Update(ctx context.Context, param model.User) (model.UserResponse, *Error) {
	var emptyUUID uuid.UUID

	hash, err := HashPassword(ctx, param.Password)
	if err != nil {
		status := http.StatusInternalServerError
		message := "error on hash User password"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "update/User",
			message, status)
	}

	param.Password = hash

	i, err := param.Update(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on update User"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "update/User",
			message, status)
	}
	return param.Response(), nil
}

func (m UserModule) Delete(ctx context.Context, param model.User) (model.UserResponse, *Error) {
	var emptyUUID uuid.UUID
	i, err := param.Delete(ctx, m.db)
	if err != nil || i == emptyUUID {
		status := http.StatusInternalServerError
		message := "error on delete User"

		return model.UserResponse{}, NewErrorWrap(err, m.Name, "delete/User",
			message, status)
	}
	return param.Response(), nil
}
