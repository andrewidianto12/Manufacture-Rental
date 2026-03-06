package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andrewidianto12/Manufacture-Rental/app/rental-manufacture/handler"
	customvalidator "github.com/andrewidianto12/Manufacture-Rental/app/rental-manufacture/validator"
	"github.com/andrewidianto12/Manufacture-Rental/database"
	equipment_repo "github.com/andrewidianto12/Manufacture-Rental/repository/equipment"
	equipment_category_repo "github.com/andrewidianto12/Manufacture-Rental/repository/equipment_category"
	maintenance_repo "github.com/andrewidianto12/Manufacture-Rental/repository/maintenance"
	notification_repo "github.com/andrewidianto12/Manufacture-Rental/repository/notification"
	payment_repo "github.com/andrewidianto12/Manufacture-Rental/repository/payment"
	rental_repo "github.com/andrewidianto12/Manufacture-Rental/repository/rental"
	report_repo "github.com/andrewidianto12/Manufacture-Rental/repository/report"
	users "github.com/andrewidianto12/Manufacture-Rental/repository/user"
	equipment_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment"
	equipment_category_service "github.com/andrewidianto12/Manufacture-Rental/service/equipment_category"
	maintenance_service "github.com/andrewidianto12/Manufacture-Rental/service/maintenance"
	notification_service "github.com/andrewidianto12/Manufacture-Rental/service/notification"
	payment_service "github.com/andrewidianto12/Manufacture-Rental/service/payment"
	rental_service "github.com/andrewidianto12/Manufacture-Rental/service/rental"
	report_service "github.com/andrewidianto12/Manufacture-Rental/service/report"
	user_service "github.com/andrewidianto12/Manufacture-Rental/service/user"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	db := database.InitDatabaseWithGORM()
	if err := db.AutoMigrate(
		&user_service.User{},
		&equipment_category_service.EquipmentCategory{},
		&equipment_service.Equipment{},
		&rental_service.Rental{},
		&payment_service.Payment{},
		&maintenance_service.Maintenance{},
		&notification_service.Notification{},
	); err != nil {
		log.Fatal(err)
	}

	// user
	userRepo := users.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// equipment
	equipmentCategoryRepository := equipment_category_repo.NewEquipmentCategoryRepository(db)
	equipmentCategoryService := equipment_category_service.NewEquipmentCategoryService(equipmentCategoryRepository)
	equipmentCategoryHandler := handler.NewEquipmentCategoryHandler(equipmentCategoryService)

	equipmentRepository := equipment_repo.NewEquipmentRepository(db)
	equipmentService := equipment_service.NewEquipmentService(equipmentRepository)
	equipmentHandler := handler.NewEquipmentHandler(equipmentService)

	// rental
	rentalRepository := rental_repo.NewRentalRepository(db)
	rentalService := rental_service.NewRentalService(rentalRepository)
	rentalHandler := handler.NewRentalHandler(rentalService)

	// payment
	paymentRepository := payment_repo.NewPaymentRepository(db)
	paymentService := payment_service.NewPaymentService(paymentRepository)
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// maintenance
	maintenanceRepository := maintenance_repo.NewMaintenanceRepository(db)
	maintenanceService := maintenance_service.NewMaintenanceService(maintenanceRepository)
	maintenanceHandler := handler.NewMaintenanceHandler(maintenanceService)

	// notification
	notificationRepository := notification_repo.NewNotificationRepository(db)
	notificationService := notification_service.NewNotificationService(notificationRepository)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	// report
	reportRepository := report_repo.NewReportRepository(db)
	reportService := report_service.NewReportService(reportRepository)
	reportHandler := handler.NewReportHandler(reportService)

	e := echo.New()
	e.Validator = customvalidator.NewValidator()

	e.GET("/swagger/openapi.yaml", func(c echo.Context) error {
		return c.File(getOpenAPIPath())
	})

	e.GET("/swagger", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<!doctype html>
<html>
	<head>
		<meta charset="UTF-8" />
		<title>Manufacture Rental API Docs</title>
		<link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@5/swagger-ui.css" />
	</head>
	<body>
		<div id="swagger-ui"></div>
		<script src="https://unpkg.com/swagger-ui-dist@5/swagger-ui-bundle.js"></script>
		<script>
			window.onload = function () {
				SwaggerUIBundle({
					url: '/swagger/openapi.yaml',
					dom_id: '#swagger-ui'
				});
			}
		</script>
	</body>
</html>`)
	})

	api := e.Group("/api")
	{
		usersGroup := api.Group("/users")
		{
			usersGroup.POST("/register", userHandler.RegisterUser)
			usersGroup.POST("/login", userHandler.LoginUser)
			usersGroup.DELETE("/:id", userHandler.DeleteUser)
		}

		equipmentCategoryGroup := api.Group("/equipment-categories")
		{
			equipmentCategoryGroup.POST("", equipmentCategoryHandler.CreateCategory)
			equipmentCategoryGroup.GET("", equipmentCategoryHandler.GetAllCategories)
			equipmentCategoryGroup.GET("/:id", equipmentCategoryHandler.GetCategoryByID)
			equipmentCategoryGroup.DELETE("/:id", equipmentCategoryHandler.DeleteCategory)
		}

		equipmentGroup := api.Group("/equipment")
		{
			equipmentGroup.POST("", equipmentHandler.CreateEquipment)
			equipmentGroup.GET("", equipmentHandler.GetAllEquipment)
			equipmentGroup.GET("/:id", equipmentHandler.GetEquipmentByID)
			equipmentGroup.DELETE("/:id", equipmentHandler.DeleteEquipment)
		}

		rentalsGroup := api.Group("/rentals")
		{
			rentalsGroup.POST("", rentalHandler.CreateRental)
			rentalsGroup.GET("", rentalHandler.GetAllRentals)
			rentalsGroup.GET("/:id", rentalHandler.GetRentalByID)
			rentalsGroup.DELETE("/:id", rentalHandler.DeleteRental)
		}

		paymentsGroup := api.Group("/payments")
		{
			paymentsGroup.POST("", paymentHandler.CreatePayment)
			paymentsGroup.GET("", paymentHandler.GetAllPayments)
			paymentsGroup.GET("/:id", paymentHandler.GetPaymentByID)
			paymentsGroup.DELETE("/:id", paymentHandler.DeletePayment)
		}

		maintenanceGroup := api.Group("/maintenance")
		{
			maintenanceGroup.POST("", maintenanceHandler.CreateMaintenance)
			maintenanceGroup.GET("", maintenanceHandler.GetAllMaintenance)
		}

		notificationsGroup := api.Group("/notifications")
		{
			notificationsGroup.POST("", notificationHandler.CreateNotification)
			notificationsGroup.GET("", notificationHandler.GetAllNotifications)
		}

		reportsGroup := api.Group("/reports")
		{
			reportsGroup.GET("/dashboard", reportHandler.GetDashboardReport)
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
	if err := godotenv.Load(".env"); err == nil {
		return
	}
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Warning: .env file not found, menggunakan sistem env")
	}
}

func getOpenAPIPath() string {
	if _, err := os.Stat("docs/openapi.yaml"); err == nil {
		return "docs/openapi.yaml"
	}

	if _, err := os.Stat("../../docs/openapi.yaml"); err == nil {
		return "../../docs/openapi.yaml"
	}

	return "docs/openapi.yaml"
}
