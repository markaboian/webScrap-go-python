package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"go-python/cmd/server/handler"
	"go-python/internal/product"
	"go-python/package/store"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error while trying to load .env file")
	}

	token := os.Getenv("TOKEN")
	userDB := os.Getenv("USER_DB")
	passwordDB := os.Getenv("PASSWORD_DB")

	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/premier_league", userDB, passwordDB)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if errPing != nil{
		panic(errPing)
	}

	storage := store.Sql{DB: db}
	repository := product.Repository{Interface: &storage}
	service := product.Service{Repository: &repository}
	handler := handler.Handler{Service: &service}

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("token")

		if tokenHeader != token {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error: invalid token",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	})

	teams := r.Group("teams")
	{
		teams.GET("/:rank", handler.GetTeamByRank)
	}

	r.Run(":8080")

}