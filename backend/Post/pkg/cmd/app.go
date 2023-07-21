package cmd

import (
	"fmt"
	"github.com/fnaf-enjoyers/post-service/pkg/config"
	"github.com/fnaf-enjoyers/post-service/pkg/repository"
	"github.com/fnaf-enjoyers/post-service/pkg/route"
	"github.com/fnaf-enjoyers/post-service/pkg/usecases"
	"github.com/gofiber/fiber/v2"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName: "Post UseCase",
	})

	err := os.Setenv("CONFIG_PATH", "./config.yml")
	if err != nil {
		log.Fatal(err)
	}

	err = os.Setenv("PORT", ":3001")
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.ReadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Fatalf("Cannot read config file: %s", err)
	}

	username := cfg.Repository.DB.Username
	password := cfg.Repository.DB.Password
	host := cfg.Repository.DB.Host
	name := cfg.Repository.DB.Name
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@%s/%s",
		username,
		password,
		host,
		name,
	)

	db, err := sqlx.Connect("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			fmt.Printf("Connection is not closed: %s", err)
		}
	}(db)

	ser := usecases.NewService(cfg)
	repo := repository.NewRepository(db)

	route.SetupSwagger(app)
	route.SetupRoutes(app, ser, repo)

	err = app.Listen(os.Getenv("PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
