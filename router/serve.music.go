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

	// f, errF := os.Open(m.FilePath)
	// if errF != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Println(errF)
	// 	return
	// }
	// defer f.Close()

	// fi, errFs := f.Stat()
	// if errFs != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Println(errFs)
	// 	return
	// }

	// w.Header().Set("Content-Type", "audio/mpeg")
	// w.Header().Set("Cache-Control", "no-cache")
	// w.Header().Set("Connection", "keep-alive")
	// w.Header().Set("Content-Length", fmt.Sprint(fi.Size()))

	// bfr := bufio.NewReader(f)
	// buf := make([]byte, 0, 4*1024)

	// for {
	// 	n, err := bfr.Read(buf[:cap(buf)])
	// 	buf = buf[:n]
	// 	if n == 0 {
	// 		if err == nil {
	// 			continue
	// 		}
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		return
	// 	}

	// 	w.Write(buf)
	// 	flusher, ok := w.(http.Flusher)
	// 	if !ok {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}
	// 	flusher.Flush()

	// 	if err != nil && err != io.EOF {
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		return
	// 	}

	// }

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Cache-Control", "no-cache")
	http.ServeFile(w, r, m.FilePath)
}
