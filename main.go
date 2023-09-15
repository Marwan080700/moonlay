package main

import (
	"fmt"
	"moonlay/database"
	"moonlay/pkg/mysql"
	"moonlay/routes"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	mysql.DatabaseInit()
	database.RunMigration()
	

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}