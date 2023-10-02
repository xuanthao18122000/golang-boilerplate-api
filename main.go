package main
import (
    "fmt"
    "log"
    "os"

    "golang-boilerplate-api/storage"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
)

type Post struct {
    Name    string `json:"name"`
	Status     string `json:"status"`
}

type Repository struct {
    DB *gorm.DB
}

func main() {
    err := godotenv.Load(".env");
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

    db, err := storage.NewConnection(config);

    if err != nil {
        log.Fatal("Connect database failed: ", err, db);
    }
    fmt.Println("App running Port 8080");
}