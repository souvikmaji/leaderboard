package main

import (
	"log"

	"github.com/bxcodec/faker/v3"
	"github.com/cheggaaa/pb/v3"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/souvikmaji/leaderboard/models"
)

func main() {
	configuration := models.InitConfig()

	log.Println("Connecting to datbase: ", configuration.Database.Name)
	db, err := db.NewDB(configuration.GetDbURI())
	if err != nil {
		log.Fatalf("failed to connect database: %s", err.Error())
	}
	defer db.Close()

	userCount := 1000

	log.Printf("Creating %d new users\n", userCount)
	bar := pb.Full.Start(userCount)

	for i := 0; i < userCount; i++ {
		user := &models.User{
			Username: faker.Username(),
			Email:    faker.Email(),
		}
		err = db.SaveUser(user)
		if err != nil {
			log.Print("Error saving user", err)
		}
		bar.Increment()
	}

	bar.Finish()
	log.Print("Users created")

}
