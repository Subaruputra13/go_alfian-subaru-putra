package routes

import (
	"praktikum/controllers"
	m "praktikum/middleware"
	"praktikum/repository/database"
	"praktikum/usecase"
	"praktikum/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"gorm.io/gorm"
)

func NewRoute(e *echo.Echo, db *gorm.DB) {
	// Middleware
	m.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// Interface
	userRepository := database.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)

	authRepository := database.NewAuthRepository(db)
	authUsecase := usecase.NewAuthUsecase(authRepository, userRepository)

	authController := controllers.NewAuthController(userUsecase, authUsecase)

	//validator
	e.Validator = &util.CustomValidator{Validator: validator.New()}

	// Routes

	e.POST("/register", authController.RegisterController)
	e.POST("/login", authController.LoginController)

}
