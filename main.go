package main

import (
	"debtomate/module/feature/middleware"
	"debtomate/module/feature/route"
	"debtomate/utils/database"
	"debtomate/utils/token"
	"debtomate/utils/viper"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {
	app := fiber.New()

	middleware.SetupMiddlewares(app)

	err := viper.ViperConfig.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}

	db, err := database.InitPGSDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	database.Migrate(db)

	jwtService := token.NewJWT(viper.ViperConfig.GetStringValue("SECRET"))

	route.SetupRoutes(app, db, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Ruti Store")
	})

	err = app.Listen(":" + port)
	if err != nil {
		log.Fatalf("Failed to start the server: %s", err)
	}
}
