package main

import (
	"errors"
	"fmt"
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
	configuration, err := models.InitConfig()
	if err != nil {
		log.Fatalln("Error initializing configuration", err)
	}

	log.Println("Connecting to datbase: ", configuration.Database.Name)

	// try connecting to database
	db, err := connectDB(configuration)
	if err != nil {
		log.Fatalln("Failed to connect database: ", err)
	}
	defer db.Close()

	log.Println("Migrating database tables")
	db.DB.AutoMigrate(&models.Game{})
	db.DB.AutoMigrate(&models.User{})
	db.DB.AutoMigrate(&models.GameUser{})

	if err := prepareData(db); err != nil {
		log.Fatalln("Failed preparing sample data:", err)
	}

	log.Println("Successfully initialized database with sample data")
}

func prepareData(db *db.DB) error {
	userCount := 1000

	_, totalUser, _, err := db.GetAllSortedGameUser(0, 0)
	if err != nil {
		return fmt.Errorf("Error fetching game user data")
	}
	if totalUser >= int64(userCount) {
		log.Printf("Total %d sample user present\n", totalUser)
		return nil
	}

	log.Println("Preparing data")
	log.Println("Creating game")

	game := &models.Game{
		Name: faker.Word(),
	}
	if err := db.SaveGame(game); err != nil {
		return fmt.Errorf("Error creating game %v", err)
	}

	log.Printf("Creating %d new users and joining game: %s\n", userCount, game.Name)
	bar := pb.Full.Start(userCount)

	for i := 0; i < userCount; i++ {
		user := &models.User{
			Username: faker.Username(),
			Email:    faker.Email(),
		}

		if err := db.SaveUser(user); err != nil {
			return fmt.Errorf("Error saving user %v", err)
		}

		if err = db.SaveGameUser(&models.GameUser{UserID: user.ID, GameID: game.ID, Score: rand.Float64() * 100}); err != nil {
			return fmt.Errorf("Error saving game user %v", err)
		}
		bar.Increment()
	}

	bar.Finish()
	log.Print("Users created")

	return nil
}

func connectDB(c *models.Configurations) (*db.DB, error) {
	store, err := db.NewDB(c)
	if err != nil {
		// if database does not exist
		var e *pq.Error
		if errors.As(err, &e) && e.Code.Name() == "invalid_catalog_name" {
			log.Println("Database does not exist")

			store, err := createDB(c)
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

func createDB(c *models.Configurations) (*db.DB, error) {
	log.Println("Creating Datbase", c.Database.Name)

	// error creating database
	if out, err := pqCreateDB(c); err != nil {
		log.Printf("Error: %v", string(out))
		return nil, fmt.Errorf("Error creating database: %v", err)
	}
	log.Println("Database created successfully")

	// try reconnecting
	log.Println("Reconnecting to datbase:", c.Database.Name)

	store, err := db.NewDB(c)
	if err != nil {
		return nil, fmt.Errorf("Error reconnecting to newly created db: %v", err)
	}

	return store, nil
}

func pqCreateDB(c *models.Configurations) ([]byte, error) {

	cmd := exec.Command("createdb", "--host", c.Database.Host, "--port", strconv.Itoa(c.Database.Port),
		"--username", c.Database.Username, "--owner", c.Database.Username, c.Database.Name, "--echo")

	if out, err := cmd.CombinedOutput(); err != nil {
		return out, err
	}

	return nil, nil
}
