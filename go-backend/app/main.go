package main

import (
	_ "database/sql"
	"go-backend/database"
	"go-backend/router"
	"go-backend/setup"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	/* Load Environment Variable */

	/* Connect to database */
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	/* Setup router */
	router.SignupRouterSetup(&setup.Env{}, db, &gin.RouterGroup{})

	/* Run */
}