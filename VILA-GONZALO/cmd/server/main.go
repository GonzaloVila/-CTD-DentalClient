package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/GonzaloVila/clinica/cmd/server/routes"
	
	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)


func main() {
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbName)

	db, err := sql.Open("mysql", dbConnection)

	if err != nil {
		panic(err)
	}
	r := gin.Default()
	
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })


	
	router := routes.NewRouter(r, db)
	router.MapRoutes()
	if err := r.Run(); err != nil {
		panic(err)
	}
}
