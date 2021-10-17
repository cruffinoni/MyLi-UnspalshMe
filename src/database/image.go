package database

import (
	"github.com/cruffinoni/MyLi-UnspalshMe/src/unsplash"
	"log"
)

func (db Database) AddImage(model unsplash.ImageModel) error {
	stmt, err := db.handler.Prepare("INSERT INTO `images` (`imageId`, `userId`) VALUES (?, ?);")
	if err != nil {
		log.Printf("can't prepare the query to add an image: %v\n", err)
		return err
	}
	if _, err = stmt.Exec(model.ID, model.User.ID); err != nil {
		log.Printf("can't execute the query: %v\n", err)
		return err
	}
	return nil
}
