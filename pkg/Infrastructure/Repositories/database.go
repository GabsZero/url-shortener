package repositories

import (
	"fmt"
	"os"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Repository struct {
}

func (r *Repository) GetDbInstance() *gorm.DB {
	return db
}

func (r *Repository) StartDabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("URL_SHORTENER_DATABASE_HOST")
	user := os.Getenv("URL_SHORTENER_DATABASE_USER")
	password := os.Getenv("URL_SHORTENER_DATABASE_PASSWORD")
	database := os.Getenv("URL_SHORTENER_DATABASE_DATABASE")
	port := os.Getenv("URL_SHORTENER_DATABASE_PORT")
	// connectionString := "root:mysqlpw@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)

	databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db = databaseConnection

	db.AutoMigrate(models.Url{})

}
