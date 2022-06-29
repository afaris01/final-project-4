package main

import (
	"final-project-4/database"
	"final-project-4/router"
)

func main() {
	database.MulaiDB()
	r := router.MulaiApp()
	r.Run(":8080")
}
