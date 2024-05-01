package main

import (
	"UsersAPI/cmd/server/handler"
	"UsersAPI/internal/product"
	"UsersAPI/package/store"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error when trying to load the file .env")
	}

	token := os.Getenv("TOKEN")
	userDB := os.Getenv("USER_DB")
	passDB := os.Getenv("PASS_DB")

	connectionString := fmt.Sprintf("%v:%v@tcp(localhost:3306)/UsersAPI", userDB, passDB)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	errPing := db.Ping()
	if errPing != nil {
		panic(errPing)
	}

	storage := store.Sql{DB: db}
	repository := product.Repository{Interface: &storage}
	service := product.Service{Repository: &repository}
	handler := handler.Handler{Service: &service}

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		tokenHeader := ctx.GetHeader("token")

		if token != tokenHeader {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Error: invalid token",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	})


	users := r.Group("users")
	{
		users.GET("/:id", handler.GetUserById)
		users.POST("/add", handler.AddUser)
		users.PATCH("/:id", handler.UpdateUser)
	}

	r.Run(":8080")
}