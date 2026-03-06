package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andrewidianto12/Manufacture-Rental/app/rental-manufacture/handler"
	customvalidator "github.com/andrewidianto12/Manufacture-Rental/app/rental-manufacture/validator"
	"github.com/andrewidianto12/Manufacture-Rental/database"
	users "github.com/andrewidianto12/Manufacture-Rental/repository/user"
	user_service "github.com/andrewidianto12/Manufacture-Rental/service/user"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	db := database.InitDatabaseWithGORM()

	// user
	userRepo := users.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e := echo.New()
	e.Validator = customvalidator.NewValidator()

	api := e.Group("/api")
	{
		usersGroup := api.Group("/users")
		{
			usersGroup.POST("/register", userHandler.RegisterUser)
			usersGroup.POST("/login", userHandler.LoginUser)
			usersGroup.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	addr := ":" + port
	fmt.Println("Connect to the database successfully (postgres)")
	log.Println("Application initialized on " + addr)
	e.Logger.Fatal(e.Start(addr))
}

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: .env file not found, menggunakan sistem env")
	}
}
