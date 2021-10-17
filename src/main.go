package main

import (
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash"
	"log"
)

func main() {
	app, err := application.New()
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	defer app.CloseDatabase()
	var reqData unsplash.SearchQueryModel
	reqData, err = app.GetApi().GetImages(app.GetArg())
	for _, i := range reqData.Results {
		if err = app.AddImageToDatabase(i); err != nil {
			return
		}
	}
}
