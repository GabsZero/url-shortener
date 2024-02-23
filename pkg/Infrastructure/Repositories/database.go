package repositories

import (
	"fmt"
	"os"

	models "github.com/gabszero/url-shortener/pkg/Infrastructure/Models"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbShards []*gorm.DB
var redisInstance *redis.Client

type Repository struct {
}

func (r *Repository) GetDbInstance(shard int) *gorm.DB {
	return dbShards[shard-1]
}

func (r *Repository) GetRedisInstance() *redis.Client {
	return redisInstance
}

func (r *Repository) StartDabase() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	count := 0
	shards := []string{"1", "2"}

	for _, shard := range shards {
		connectionString := GetConnectionString(shard)
		databaseConnection, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		databaseConnection.AutoMigrate(models.Url{})
		dbShards = append(dbShards, databaseConnection)

		count++
	}

	//initializing redis
	ri := redis.NewClient(&redis.Options{
		Addr:     "cache:6379",
		Password: "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81",
		DB:       0, // use default DB
	})

	redisInstance = ri

}

func GetConnectionString(shard string) string {
	host := os.Getenv("URL_SHORTENER_DATABASE_HOST_" + shard)
	user := os.Getenv("URL_SHORTENER_DATABASE_USER_" + shard)
	password := os.Getenv("URL_SHORTENER_DATABASE_PASSWORD_" + shard)
	database := os.Getenv("URL_SHORTENER_DATABASE_DATABASE_" + shard)
	port := os.Getenv("URL_SHORTENER_DATABASE_PORT_" + shard)
	// connectionString := "root:mysqlpw@tcp(127.0.0.1:3306)/url_shortener?charset=utf8mb4&parseTime=True&loc=Local"
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, database)
}
