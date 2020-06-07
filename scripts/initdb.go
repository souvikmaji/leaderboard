package main

import (
	"log"
	"math/rand"

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

	log.Println("Migrating database tables")
	db.DB.AutoMigrate(&models.Game{})
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.GameUser{})

	log.Println("Creating game")
	game := &models.Game{
		Name: faker.Word(),
	}
	db.SaveGame(game)

	userCount := 1000

	log.Printf("Creating %d new users and joining %s\n", userCount, game.Name)
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

		db.SaveGameUser(&models.GameUser{UserID: user.ID, GameID: game.ID, Score: rand.Float64() * 100})
		bar.Increment()
	}

	bar.Finish()
	log.Print("Users created")

}
