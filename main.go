package main

import (
	"ProductSearchAppSql/config"
	"ProductSearchAppSql/routes"
	"fmt"
	"net/http"
)

func main() {
	config.ConnectDB()
	// config.AutoMigrate()

	r := routes.Router()
	fmt.Println("Server started at localhost:3000")
	http.ListenAndServe(":3000", r)
} 