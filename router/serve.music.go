package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/renosyah/simple-podcast/model"
	uuid "github.com/satori/go.uuid"
)

func HandlerFileMusic(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)

	id, err := uuid.FromString(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	m, errM := musicModule.One(ctx, model.Music{ID: id})
	if errM != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(errM)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	http.ServeFile(w, r, m.FilePath)
}
