package routes

import (
	"praktikum/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	// User Controller
	e.GET("/users", controllers.GetUserControllers)
	e.POST("/users", controllers.CreateUserController)
	e.GET("/users/:id", controllers.GetUserByIdController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)

	// Book Controller
	e.GET("/books", controllers.GetBookController)
	e.POST("/books", controllers.CreateBookController)
	e.GET("/books/:id", controllers.GetBookByIdController)
	e.PUT("/books/:id", controllers.UpdateBookByIdController)
	e.DELETE("/books/:id", controllers.DeleteBookByIdController)

	return e
}
