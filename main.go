package main

import (
	"database/sql"
	"fmt"
	"formative-15/controllers"
	"formative-15/database"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	//ENV
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed to load environment")
	} else {
		fmt.Println("success read file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	DB, err = sql.Open("postgres", psqlInfo)
	err = DB.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(DB)

	defer DB.Close()

	// Router GIN

	r := gin.Default()
	r.GET("/persons", controllers.GetAllPerson)
	r.POST("/persons", controllers.InsertPerson)
	r.PUT("/persons/:id", controllers.UpdatePerson)
	r.DELETE("/persons/:id", controllers.DeletePerson)

	r.Run(":8080")
}
