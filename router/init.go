package router

import (
	"context"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/pkg/errors"
	"github.com/renosyah/simple-podcast/api"
)

var (
	dbPool *sql.DB

	userModule     *api.UserModule
	categoryModule *api.CategoryModule
	musicModule    *api.MusicModule
	podcastModule  *api.PodcastModule
)

func Init(db *sql.DB) {
	dbPool = db

	userModule = api.NewUserModule(db)
	categoryModule = api.NewCategoryModule(db)
	musicModule = api.NewMusicModule(db)
	podcastModule = api.NewPodcastModule(db)
}

func ParseBodyData(ctx context.Context, r *http.Request, data interface{}) error {
	bBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return errors.Wrap(err, "read")
	}

	err = json.Unmarshal(bBody, data)
	if err != nil {
		return errors.Wrap(err, "json")
	}

	valid, err := govalidator.ValidateStruct(data)
	if err != nil {
		return errors.Wrap(err, "validate")
	}

	if !valid {
		return errors.New("invalid data")
	}

	return nil
}
