package main

import (
	"github.com/Thomika1/TestOrga/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.InitializeRoutes(r)

	r.Run()

} // function main
