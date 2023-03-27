package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User
var mapUsers map[int]User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	userId := getIdFromParam(c)

	if val, ok := mapUsers[userId]; ok {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"messages": "success get by id users",
			"user":     val,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "error get user by id",
		"users":    nil,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	userId := getIdFromParam(c)

	for index, val := range users {
		if val.Id == userId {
			users = append(users[:index], users[index+1:]...)
			delete(mapUsers, userId)

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete by id users",
			})
		}

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "error delete user by id",
		"users":    nil,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	userId := getIdFromParam(c)

	for index, val := range users {
		if val.Id == userId {
			user := User{}
			c.Bind(&user)
			users[index] = user
			mapUsers[userId] = user

			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update by id users",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "error update user by id",
		"users":    nil,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
		mapUsers = make(map[int]User)
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	mapUsers[user.Id] = user
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// function for get id from param
func getIdFromParam(c echo.Context) int {
	id, _ := strconv.Atoi(c.Param("id"))
	return id
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.GET("/users/:id", GetUserController)
	e.POST("/users", CreateUserController)
	e.PUT("/users/:id", UpdateUserController)
	e.DELETE("/users/:id", DeleteUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
