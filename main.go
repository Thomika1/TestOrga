package main

import (
	"github.com/Thomika1/TestOrga/database"
	"github.com/Thomika1/TestOrga/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	dbConnection, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	routes.InitializeRoutes(r)

	r.Run(":8080")

} // function main
