package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"syscall/js"

	"github.com/Magnific86/pg-golang/server/models"
	"github.com/Magnific86/pg-golang/server/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

type Blob struct {
	object *js.Value
}

type FormFile struct {
	Blob
	cur        int64
	buffersize int64
	size       int64
}

type Post struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	File    FormFile `json:"file"`
}

func (r *Repository) CreatePost(context *fiber.Ctx) error {

	postModel := new(Post)

	file, errorFile := context.FormFile("file")

	if errorFile != nil {
		log.Fatal("failed to parse file")
	}

	err := context.BodyParser(postModel)

	if err != nil {
		log.Fatal("failed to parse file")
	}

	postToPatch := Post{
		Title:   "my title",
		Content: "my content",
		File:    *file,
	}

	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	err = r.DB.Create(&postToPatch).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create post"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Post has been added"})
	return nil
}

func (r *Repository) DeletePost(context *fiber.Ctx) error {
	postModel := models.Posts{}
	id := context.Params("id")
	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	err := r.DB.Delete(postModel, id)

	if err.Error != nil {
		context.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "could not delete Post",
		})
		return err.Error
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Post delete successfully",
	})
	return nil
}

func (r *Repository) GetPosts(context *fiber.Ctx) error {
	posts := &[]models.Posts{}

	err := r.DB.Find(posts).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get Posts"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"data": posts,
	})
	return nil
}

func (r *Repository) GetPost(context *fiber.Ctx) error {
	id := context.Params("id")

	if id == "" {
		context.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "id cannot be empty",
		})
		return nil
	}

	fmt.Println("the ID is", id)

	postModel := &models.Posts{}

	err := r.DB.Where("id = ?", id).First(postModel).Error

	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get the Post"})
		return err
	}

	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Post id fetched successfully",
		"data":    postModel,
	})
	return nil
}

func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_post", r.CreatePost)
	api.Delete("delete_post/:id", r.DeletePost)
	api.Get("/posts/:id", r.GetPost)
	api.Get("/posts", r.GetPosts)
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("failed to load env vars", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
		DBName:   os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {
		log.Fatal("failed to set db connection", err)
	}

	err = models.MigratePosts(db)

	if err != nil {
		log.Fatal("failed to migrate db")
	}

	r := Repository{
		DB: db,
	}

	app := fiber.New()

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	r.SetupRoutes(app)

	app.Listen(":8080")
}
