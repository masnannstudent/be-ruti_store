package main

import (
	"github.com/gofiber/fiber/v2"
	"os"
	"ruti-store/config"
	"ruti-store/module/feature/middleware"
	"ruti-store/module/feature/route"
	"ruti-store/module/feature/user/repository"
	"ruti-store/module/feature/user/service"
	"ruti-store/utils/database"
	"ruti-store/utils/payment"
	"ruti-store/utils/token"
)

func main() {
	app := fiber.New()
	var initConfig = config.InitConfig()
	jwtService := token.NewJWT(initConfig.Secret)

	middleware.SetupMiddlewares(app)
	db := database.InitPGSDatabase(*initConfig)
	snapClient := payment.InitSnapMidtrans(*initConfig)
	coreClient := payment.InitCoreMidtrans(*initConfig)
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)

	database.Migrate(db)
	route.SetupRoutes(app, db, jwtService, snapClient, userService, coreClient)

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
