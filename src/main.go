package main

import (
	"fmt"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/application/args"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash/models"
	"log"
)

func main() {
	fmt.Print("Application initialization...\n")
	app, err := application.New()
	if err != nil {
		if err == args.HelpMessage {
			return
		}
		log.Printf("Error: %v\n", err)
		return
	}
	defer app.CloseDatabase()
	var reqData models.SearchImageQuery
	fmt.Print("Looking for your images on Unsplash API...\n")
	reqData, err = app.GetApi().GetImages(app.GetArg())
	if err != nil {
		log.Printf("could not get the images: %v\n", err)
		return
	}
	fmt.Printf("%v image(s) found", len(reqData.Results))
	for _, i := range reqData.Results {
		if err = app.AddImageToDatabase(i); err != nil {
			log.Printf("can't add image to the database: %v\n", err)
			return
		}
		description := i.Description
		if len(i.Description) == 0 {
			description = "No description provided"
		}
		fmt.Printf("- Image posted by %v:\nDescription: %v\n", i.User.Username, description)
		fmt.Printf("Image available at %v\n", i.Urls.Full)
	}
}
