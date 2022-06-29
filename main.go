package main

import (
	"final-project-4/database"
	"final-project-4/router"
	"os"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	port := os.Getenv("PORT")
	r.Run(":" + port)
}
