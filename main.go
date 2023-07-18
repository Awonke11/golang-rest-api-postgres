package main

import (
	"golang-rest-api/models"
	"golang-rest-api/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Users struct {
	user_id int    `json:"user_id`
	name    string `json:"name"`
	age     int    `json:"age"`
	phone   string `json:"phone"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetUsers(context *fiber.Ctx) error {
	users := &[]models.Users{}

	err := r.DB.Find(users).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not get users"})
		return err
	}

	context.Status(http.StatusOK).JSON(
		&fiber.Map{
			"status_code": 200,
			"data":        users,
		})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Get("/users", r.GetUsers)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal(err)
	}
	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASSWORD"),
		User:     os.Getenv("DB_USER"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Could not load the database")
	}

	err = models.MigrateUsers(db)
	if err != nil {
		log.Fatal("Could not migrate db")
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}
