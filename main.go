package main

import (
	"my-gram/database"
	"my-gram/router"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	database.StartDB()
	r := router.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
