package main

import (
	"github.com/Thomika1/TestOrga/database"
	"github.com/Thomika1/TestOrga/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}
func main() {

	r := gin.Default()

	dbConnection, err := database.ConnectDB()

	if err != nil {
		panic(err)
	}

	routes.InitializeRoutes(r, dbConnection)

	r.Run(":8080")

} // function main
