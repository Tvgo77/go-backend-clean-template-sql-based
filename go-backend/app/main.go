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
	passwd := env.DBpassword
	dsn := "host=localhost user=postgres dbname=forumdb port=5432 sslmode=disable password=" + passwd
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