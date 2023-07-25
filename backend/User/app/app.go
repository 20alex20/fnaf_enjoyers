package app

import (
	"fmt"
	"github.com/fnaf-enjoyers/user/config"
	"github.com/fnaf-enjoyers/user/repository"
	"github.com/fnaf-enjoyers/user/route"
	"github.com/fnaf-enjoyers/user/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v2"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
)

func Run() {
	app := fiber.New(fiber.Config{
		AppName: "User Service",
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	err := os.Setenv("CONFIG_PATH", "./config.yml")
	if err != nil {
		logrus.Fatal(err)
	}

	err = os.Setenv("PORT", ":3002")
	if err != nil {
		logrus.Fatal(err)
	}

	cfg, err := config.ReadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		logrus.Fatalf("Cannot read config file: %s", err)
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
		logrus.Fatal(err)
	}
	defer func(db *sqlx.DB) {
		err = db.Close()
		if err != nil {
			fmt.Printf("Connection is not closed: %s", err)
		}
	}(db)

	config.Store = session.New(session.Config{
		Storage: postgres.New(postgres.Config{
			ConnectionURI: connStr}),
	})

	uc := usecase.NewService(cfg)
	repo := repository.NewRepository(db)

	route.SetupSwagger(app)
	route.SetupRoutes(app, uc, repo)

	err = app.Listen(os.Getenv("PORT"))
	if err != nil {
		logrus.Fatal(err)
	}
}
