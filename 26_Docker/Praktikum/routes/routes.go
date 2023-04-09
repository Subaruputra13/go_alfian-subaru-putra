package routes

import (
	"praktikum/constants"
	"praktikum/controllers"
	"praktikum/middleware"

	mid "github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	// Crate echo intance
	e := echo.New()

	// Midlleware
	middleware.LogMiddleware(e)
	e.Pre(mid.RemoveTrailingSlash())

	// User Controller
	us := e.Group("/users")
	us.GET("", controllers.GetUserControllers, mid.JWT([]byte(constants.SCREAT_JWT)))
	us.POST("", controllers.CreateUserController)
	us.GET("/:id", controllers.GetUserByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))
	us.PUT("/:id", controllers.UpdateUserByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))
	us.DELETE("/:id", controllers.DeleteUserByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))
	us.POST("/login", controllers.LoginUserController)

	// Book Controller
	bk := e.Group("/books")
	bk.GET("", controllers.GetBookController, mid.JWT([]byte(constants.SCREAT_JWT)))
	bk.POST("", controllers.CreateBookController, mid.JWT([]byte(constants.SCREAT_JWT)))
	bk.GET("/:id", controllers.GetBookByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))
	bk.PUT("/:id", controllers.UpdateBookByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))
	bk.DELETE("/:id", controllers.DeleteBookByIdController, mid.JWT([]byte(constants.SCREAT_JWT)))

	// Blog Controller
	bl := e.Group("/blogs")
	bl.GET("", controllers.GetBlogController)
	bl.POST("", controllers.CreateBlogController)
	bl.GET("/:id", controllers.GetBlogByIdController)
	bl.PUT("/:id", controllers.UpdateBlogByIdController)
	bl.DELETE("/:id", controllers.DeleteBlogByIdController)

	//Auth Basic
	// eAuthBasic := e.Group("/auth")
	// eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	// eAuthBasic.GET("/users", controllers.GetUserControllers)

	return e
}
