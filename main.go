package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"ruti-store/config"
	"ruti-store/module/feature/middleware"
	"ruti-store/module/feature/route"
	"ruti-store/utils/database"
	"ruti-store/utils/token"
)

func main() {
	app := fiber.New()
	var initConfig = config.InitConfig()
	jwtService := token.NewJWT(initConfig.Secret)

	middleware.SetupMiddlewares(app)
	db := database.InitPGSDatabase(*initConfig)
	database.Migrate(db)
	route.SetupRoutes(app, db, jwtService)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Ruti Store")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	err := app.Listen(":" + port)
	if err != nil {
		panic("Failed to start the server: " + err.Error())
	}

}
