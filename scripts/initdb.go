package main

import (
	"errors"
	"log"
	"math/rand"
	"os/exec"
	"strconv"

	"github.com/bxcodec/faker/v3"
	"github.com/cheggaaa/pb/v3"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/lib/pq"
	"github.com/souvikmaji/leaderboard/db"
	"github.com/souvikmaji/leaderboard/models"
)

func main() {
	configuration := models.InitConfig()
	log.Println("Connecting to datbase: ", configuration.Database.Name)

	// try connecting to database
	db, err := connectDB(configuration)
	if err != nil {
		log.Fatalln("failed to connect database: ", err)
	}
	defer db.Close()

	log.Println("Migrating database tables")
	db.DB.AutoMigrate(&models.Game{})
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.GameUser{})

	prepareData(db)

	log.Println("Successfully initialized database with sample data")
}

func prepareData(db *db.DB) {
	log.Println("Preparing data")
	log.Println("Creating game")

	game := &models.Game{
		Name: faker.Word(),
	}
	db.SaveGame(game)

	userCount := 1000

	log.Printf("Creating %d new users and joining game: %s\n", userCount, game.Name)
	bar := pb.Full.Start(userCount)

	for i := 0; i < userCount; i++ {
		user := &models.User{
			Username: faker.Username(),
			Email:    faker.Email(),
		}
		err := db.SaveUser(user)
		if err != nil {
			log.Print("Error saving user", err)
		}

		db.SaveGameUser(&models.GameUser{UserID: user.ID, GameID: game.ID, Score: rand.Float64() * 100})
		bar.Increment()
	}

	bar.Finish()
	log.Print("Users created")
}

func connectDB(c *models.Configurations) (*db.DB, error) {
	store, err := db.NewDB(c.GetDbURI())
	if err != nil {
		// if database does not exist
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "invalid_catalog_name" {
			log.Println("Database does not exist")
			out, err := createDB(c)

			// error creating database
			if err != nil {
				log.Printf("Error: %v", err)
				return nil, errors.New(string(out))
			}
			log.Println("Database created successfully")

			// try reconnecting
			log.Println("Reconnecting to datbase: ", c.Database.Name)

			store, err = db.NewDB(c.GetDbURI())
			if err != nil {
				return nil, err
			}
			return store, nil
		}

		// fail on errors other than db exist
		return nil, err
	}

	return store, nil
}

func createDB(c *models.Configurations) ([]byte, error) {
	log.Println("Creating Datbase", c.Database.Name)

	cmd := exec.Command("createdb", "--host", c.Database.Host, "--port", strconv.Itoa(c.Database.Port),
		"--username", c.Database.Username, "--owner", c.Database.Username, c.Database.Name, "--echo")

	if out, err := cmd.CombinedOutput(); err != nil {
		return out, err
	}

	return nil, nil
}