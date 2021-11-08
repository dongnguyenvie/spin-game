package main

import (
	"fmt"
	"log"
	"nolan/spin-game/components/appctx"
	"nolan/spin-game/components/middleware"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
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

	ethClient, err := ethclient.Dial(InfureEndpoint)
	if err != nil {
		fmt.Println(err)
	}

	appCtx := appctx.NewAppContext(db, ethClient, SecretKey, SmartContractAddr)

	setupInfura(appCtx)

	r := gin.Default()

	{
		r.Use(middleware.Recover(appCtx))
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	{
		v1 := r.Group("/v1")
		setupMainRoute(appCtx, v1)
	}

	r.Run(":" + port)

}
