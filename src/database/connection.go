package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Database struct {
	handler *sql.DB
}

func (db Database) createAndSelectDatabase(databaseName string) error {
	var err error
	if _, err = db.handler.Exec("CREATE DATABASE IF NOT EXISTS " + databaseName + ";"); err != nil {
		return err
	}
	_, err = db.handler.Exec("USE " + databaseName + ";")
	return err
}

func (db Database) createImageTable() error {
	const query = "CREATE TABLE IF NOT EXISTS `images` (`id` INT NOT NULL AUTO_INCREMENT, `imageId` VARCHAR(255) NOT NULL, `userId` VARCHAR(255) NOT NULL, PRIMARY KEY (`id`));"
	_, err := db.handler.Exec(query)
	return err
}

func (db Database) Close() {
	fmt.Printf("Closing database...\n")
	err := db.handler.Close()
	if err != nil {
		log.Printf("error while closing the database: %v\n", err)
	} else {
		fmt.Printf("Database successfully closed\n")
	}
}

func New() (*Database, error) {
	var (
		db            Database
		err           error
		seekedEnvKeys = []string{"MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_DB_NAME"}
		envKeys       []string
	)
	for idx, key := range seekedEnvKeys {
		envKeys = append(envKeys, os.Getenv(key))
		if len(envKeys[idx]) == 0 {
			return nil, errors.New(fmt.Sprintf("envkey %v is missing\n", key))
		}
	}
	db.handler, err = sql.Open("mysql", fmt.Sprintf("%v:%v@/", envKeys[0], envKeys[1]))
	if err != nil {
		log.Printf("unable to connect to mysql database: %v\n", err)
		db.Close()
		return nil, err
	}
	if err = db.createAndSelectDatabase(envKeys[2]); err != nil {
		log.Printf("Can't create or select the database: %v\n", err)
		db.Close()
		return nil, err
	}
	if err = db.createImageTable(); err != nil {
		log.Printf("Can't create the image table")
		db.Close()
		return nil, err
	}
	return &db, nil
}
