package routes

import (
	"praktikum/controllers"
	"praktikum/middleware"

	"github.com/labstack/echo"
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

	// Blog Controller
	e.GET("/blogs", controllers.GetBlogController)
	e.POST("/blogs", controllers.CreateBlogController)
	e.GET("/blogs/:id", controllers.GetBlogByIdController)
	e.PUT("/blogs/:id", controllers.UpdateBlogByIdController)
	e.DELETE("/blogs/:id", controllers.DeleteBlogByIdController)

	// Midlleware
	middleware.LogMiddleware(e)

	//Auth Basic
	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	// eAuthBasic.GET("/users", controllers.GetUserControllers)
	return e
}
