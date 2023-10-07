package main

import (
	"fmt"
	"golang-boilerplate-api/models"
	"golang-boilerplate-api/storage"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Post struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type Repository struct {
	DB *gorm.DB
}

func (r *Repository) GetListPost(context *fiber.Ctx) error {
	postModels := &[]models.Posts{}

	err := r.DB.Find(postModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Could not get posts!",
		})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get list posts successful!",
		"data":    postModels,
	})
	return nil
}

func (r *Repository) GetPostByID(context *fiber.Ctx) error {
	id := context.Params("id")
	postModels := &models.Posts{}
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	err := r.DB.Where("id = ?", id).First(postModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "ID Post Not Found!",
		})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Get post by ID successful!",
		"data":    postModels,
	})
	return nil
}

func (r *Repository) CreatePost(context *fiber.Ctx) error {
	post := Post{}

	err := context.BodyParser(&post)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(&fiber.Map{
			"message": "request failed",
		})
	}

	err = r.DB.Create(&post).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "Can not create Post!",
		})
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{"message": "Create Post sucessful!"})
	return nil
}

func (r *Repository) UpdatePost(context *fiber.Ctx) error {
	return nil
}

func (r *Repository) DeletePost(context *fiber.Ctx) error {
	id := context.Params("id")
	postModels := models.Posts
	
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	bookAPI := api.Group("/posts")
	bookAPI.Get("/", r.GetListPost)
	bookAPI.Get("/:id", r.GetPostByID)
	bookAPI.Post("/", r.CreatePost)
	bookAPI.Put("/", r.UpdatePost)
	bookAPI.Delete("/", r.DeletePost)

	userAPI := api.Group("/user")
	userAPI.Get("/", r.GetListPost)
	userAPI.Get("/:id", r.GetPostByID)
	userAPI.Post("/", r.CreatePost)
	userAPI.Put("/", r.UpdatePost)
	userAPI.Delete("/", r.DeletePost)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Load env failed: ", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Password: os.Getenv("DB_PASS"),
		User:     os.Getenv("DB_USER"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DbName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("Connect database failed: ", err, db)
	}

	err = models.MigratePosts(db)
	if err != nil {
		log.Fatal("Migrate database failed: ", err, db)
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":3000")
}
