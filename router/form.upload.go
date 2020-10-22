package router

import (
	"fmt"
	"io"
	"net/http"
	"os"

	uuid "github.com/satori/go.uuid"

	"github.com/pkg/errors"
	"github.com/renosyah/simple-podcast/api"
	"github.com/renosyah/simple-podcast/model"
)

func HandlerUploadFile(w http.ResponseWriter, r *http.Request) (interface{}, *api.Error) {

	var response model.UploadResponse

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Upload"),
			http.StatusText(http.StatusInsufficientStorage), http.StatusInsufficientStorage)
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Upload"),
			http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
	defer file.Close()

	path := fmt.Sprintf("files/%s/", uuid.NewV1().String())
	err = os.MkdirAll(path, 0777)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Upload"),
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	tempFile, err := os.Create(fmt.Sprintf("%s%s", path, handler.Filename))
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Upload"),
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
	defer tempFile.Close()

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return nil, api.NewError(errors.Wrap(err, "Upload"),
			http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	response = model.UploadResponse{
		Name: handler.Filename,
		Path: tempFile.Name(),
	}

	return response, nil
}
