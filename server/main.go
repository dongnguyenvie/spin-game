package main

import (
	"fmt"
	"log"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/middleware"
	"nolan/spin-game/components/pubsub"
	local_pubsub "nolan/spin-game/components/pubsub/localpb"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DatabaseHost := os.Getenv("DATABASE_HOST")
	DatabaseUsername := os.Getenv("DATABASE_USERNAME")
	DatabasePassword := os.Getenv("DATABASE_PASSWORD")
	DatabaseName := os.Getenv("DATABASE_NAME")
	DatabasePort := os.Getenv("DATABASE_PORT")
	SecretKey := os.Getenv("SECRET_KEY")
	InfureEndpoint := os.Getenv("INFURA_ENDPOINT")
	SmartContractAddr := os.Getenv("CONTRACT_ADDRESS")
	port := os.Getenv("PORT")

	dsn := DatabaseUsername + ":" + DatabasePassword + "@tcp(" + DatabaseHost + ":" + DatabasePort + ")/" + DatabaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	db = db.Debug()

	ethClient, err := ethclient.Dial(InfureEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	var localPubsub pubsub.PubSub = local_pubsub.NewPubSub()

	appCtx := appctx.NewAppContext(db, ethClient, localPubsub, SecretKey, SmartContractAddr)

	{
		setupInfura(appCtx)
	}

	{
		setupPubsubRoute(appCtx)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	{
		r.Use(middleware.Recover(appCtx))
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	{
		v1 := r.Group("/api/v1")
		setupMainRoute(appCtx, v1)
	}

	r.Run(":" + port)

}
