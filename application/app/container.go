package app

import (
	"github.com/labstack/echo/v4"
	"log"
	"project/application/controller"
	"project/application/core/entities"
	"project/application/repositories"
	"project/application/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Start() {
	// Usando una base de datos en memoria para simplificar
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&entities.User{}, &entities.Challenge{}, &entities.Video{})

	userRepo := repositories.NewUserRepository(db)
	challengeRepo := repositories.NewChallengeRepository(db)
	videoRepo := repositories.NewVideoRepository(db)

	userService := services.NewUserService(userRepo)
	challengeService := services.NewChallengeService(challengeRepo, userRepo)
	videoService := services.NewVideoService(videoRepo, userRepo)
	gptFillService := services.NewGPTFillService(userRepo, challengeRepo, videoRepo)

	e := echo.New()
	controller.RegisterUserRoutes(e, userService)
	controller.RegisterChallengeRoutes(e, challengeService)
	controller.RegisterVideoRoutes(e, videoService)
	controller.RegisterGPTFillRoutes(e, gptFillService)

	e.Logger.Fatal(e.Start(":8080"))
}
