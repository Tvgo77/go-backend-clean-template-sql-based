package main

import (
	_ "database/sql"
	"go-backend/database"
	"go-backend/router"
	"go-backend/setup"
	"go-backend/domain"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	/* Load Environment Variable */
	env := setup.NewEnv()

	/* Connect to database */
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := database.NewDatabase(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}

	/* Run database migration if set in env */
	if env.RunMigration {
		db.AutoMigrate(&domain.User{})
	}
	
	/* Setup router */
	ginEngine := gin.Default()
	publicRouter := ginEngine.Group("")

	router.SignupRouterSetup(env, db, publicRouter)

	/* Run */
}