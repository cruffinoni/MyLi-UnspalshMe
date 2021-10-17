package main

import (
	"fmt"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/database"
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash"
	"log"
	"time"
)

func main() {
	start := time.Now()
	db, err := database.New()
	if err != nil {
		log.Fatalf("can't instantiate the database: %v\n", err)
	}
	defer db.Close()
	unsplashApi := unsplash.New()
	var reqData unsplash.SearchQueryModel
	reqData, err = unsplashApi.GetImages("https://api.unsplash.com/search/photos?query=dog")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return
	}
	for _, i := range reqData.Results {
		if err = db.AddImage(i); err != nil {
			return
		}
	}
	fmt.Printf("Added %v images in %v ms\n", len(reqData.Results), (time.Now().UnixNano()-start.UnixNano())/int64(time.Millisecond))
}
