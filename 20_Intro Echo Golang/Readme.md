### RANGKUMAN INTRO ECHO GOLANG

<details open>
<summary>1. APA DAN KENAPA ECHO</summary>
<br>

**Echo** adalah High performance extensible minimalis / Web Framework dari Golang. **Keuntungan Menggunakan Echo** adalah Optimazed Router, Data Rendering, Data Binding, Middleware, Scalable

</details>

<details>
<summary>2. INSTALL ECHO</summary>
<br>

Buka Terminal dan tuliskan script ini :
`go get github.com/labstack/echo/v4`

</details>

<details>
<summary>3. BASIC ROUTES & CONTROLLER</summary>
<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func HelloController(c echo.Context) error {
	return c.String(http.StatusOk, "Hello World")
}

func main() {
	// Membuat echo instance
	e := echo.Now()
	// Membuat Route ke function
	e.GET("/", HelloControler)
	// Memulai Server
	e.START(":8000")
}
```

</details>

<details>
<summary>4. RENDER DATA (JSON RESPONSE)</summary>
<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User Struct {
	Name  string
	Email string
}

func GetUser(c echo.Context) error {
	user := User{
		Name : "Ismaol",
		Email : "ismail@gmail.com"
	}

	return c.JSON(http.StatusOk, user)
}

func main() {
	e := echo.Now()
	e.GET("/", GetUser)
	e.START(":8000")
}
```

</details>

<details>
<summary>5. RETRIEVE DATA</summary>
<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User struct {
	Id       int
	Name     string
	Email    string
}

func GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user := User{
		Name : "Ismail",
		Email : "ismail@gmail.com"
	}
	// Render Data - JSON Response
	return c.JSON(http.StatusOk, map[string]interface{}{
		"user" : user,
	})
}

func main() {
	e := echo.Now()
	e.GET("/user/:id", GetUserController)
	e.START(":8000")
}
```

<br>

### Query Params

<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func UserSearchController(c echo.Context) error {
	match := c.QueryParam("match")
	return c.JSON(http.StatusOk, map[string]interface{}{
		"match" : match,
		"result" : []string {"adi","aan","asif"}, // hardcode data
	})
}

func main() {
	e := echo.Now()
	// Routing dengan Query Parameter
	e.GET("/users", UserSearchController)
	e.START(":8000")
}
```

<br>

### Form Value

<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User Struct {
	Name  string `json:"name" form: "name"`
	Email string `json:"email" form: "email"`
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	var user User
	user.Name = name
	user.Email = email

	return c.JSON(http.StatusOk, map[string]interface{}{
		"message" : "succes create user",
		"user"    : user,
	})
}

func main() {
	e := echo.Now()
	e.GET("/users", CreateUser)
	e.START(":8000")
}
```

</details>

<details>
<summary>6. BINDING DATA</summary>
<br>

```go
package main

import (
	"net/http"

	"github.com/labstack/echo"
)

type User Struct {
	Name  string `json:"name" form: "name"`
	Email string `json:"email" form: "email"`
}

func CreateUser(c echo.Context) error {
	//binding data
	user := user{}
	c.Bind(&user)

	return c.JSON(http.StatusOk, map[string]interface{}{
		"message" : "succes create user",
		"user"    : user,
	})
}

func main() {
	e := echo.Now()
	e.GET("/users", CreateUser)
	e.START(":8000")
}
```

</details>
