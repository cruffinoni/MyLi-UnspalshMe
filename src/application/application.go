package application

import (
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application/args"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/database"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash/models"
)

type Application struct {
	database *database.Database
	args     args.ProgArg
	api      unsplash.Api
}

func (app Application) CloseDatabase() {
	app.database.Close()
}

func (app Application) AddImageToDatabase(image models.Image) error {
	return app.database.AddImage(image)
}

func (app Application) GetApi() unsplash.Api {
	return app.api
}

func (app Application) GetArg() args.ProgArg {
	return app.args
}

func New() (*Application, error) {
	var (
		err error
		app Application
	)
	if app.args, err = args.New(); err != nil {
		return nil, err
	}
	if app.api, err = unsplash.New(); err != nil {
		return nil, err
	}
	if app.database, err = database.New(); err != nil {
		return nil, err
	}
	return &app, nil
}
