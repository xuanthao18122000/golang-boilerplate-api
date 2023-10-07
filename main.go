package main

import (
	controller "golang-boilerplate-api/controllers"
	_ "golang-boilerplate-api/docs"
	models "golang-boilerplate-api/models"
	repository "golang-boilerplate-api/repositories"
	routers "golang-boilerplate-api/routers"
	service "golang-boilerplate-api/services"
	"golang-boilerplate-api/storage"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"

	"github.com/joho/godotenv"
)

type MinIOConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	Bucket     string
	SecureMode bool
}

func loadMinIOConfig() MinIOConfig {
	return MinIOConfig{
		Endpoint:   os.Getenv("MINIO_ENDPOINT"),
		AccessKey:  os.Getenv("MINIO_ACCESS_KEY"),
		SecretKey:  os.Getenv("MINIO_SECRET_KEY"),
		Bucket:     os.Getenv("MINIO_BUCKET"),
		SecureMode: false, // Đặt giá trị này là true nếu bạn sử dụng HTTPS
	}
}

// func (r *Repository) UploadFile(c *fiber.Ctx) error {
// 	minioConfig := loadMinIOConfig()

// 	// Khởi tạo client MinIO
// 	minioClient, err := minio.New(minioConfig.Endpoint, &minio.Options{
// 		Creds:  credentials.NewStaticV4(minioConfig.AccessKey, minioConfig.SecretKey, ""),
// 		Secure: false, // Nếu sử dụng HTTPS, đặt giá trị này là true
// 	})

// 	file, err := c.FormFile("file")
// 	if err != nil {
// 		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
// 			"message": "File upload failed",
// 		})
// 	}

// 	// Đọc dữ liệu từ tệp hình ảnh
// 	fileBytes, err := file.Open()
// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "File read failed",
// 		})
// 	}
// 	defer fileBytes.Close()

// 	// Tạo tên tệp duy nhất
// 	objectName := fmt.Sprintf("%s/%s", minioConfig.Bucket, file.Filename)

// 	// Tải lên tệp hình ảnh lên MinIO
// 	_, err = minioClient.PutObject(c.Context(), minioConfig.Bucket, objectName, fileBytes, file.Size, minio.PutObjectOptions{
// 		ContentType: file.Header.Get("Content-Type"),
// 	})

// 	if err != nil {
// 		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
// 			"message": "File upload to MinIO failed",
// 		})
// 	}

// 	return c.Status(http.StatusOK).JSON(fiber.Map{
// 		"message": "File uploaded successfully",
// 	})
// }

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
	validate := validator.New()

	//Repository
	postsRepository := repository.NewPostsREpositoryImpl(db)

	// Service
	postsService := service.NewPostServiceImpl(postsRepository, validate)

	// Controller
	postsController := controller.NewPostsController(postsService)

	routes := routers.NewRouter(postsController)

	server := &http.Server{
		Addr:    ":3000",
		Handler: routes,
	}

	err = server.ListenAndServe()
	if err != nil {
		// Xử lý lỗi ở đây
		log.Fatalf("Server error: %v", err)
	}
}
