# RANGKUMAN MIDDLEWARE

<details open>
<summary>1. APA ITU MIDDLEWARE</summary>
<br>

**Middleware** adalah entitas yang terhubung ke pemrosesan permintaan/respons server. Middleware memungkinkan kita untuk mengimplementasikan berbagai fungsi di sekitar permintaan HTTP yang masuk ke server dan respons keluar.

</details>

<details>
<summary>2. IMPLEMENTASI MIDDLEWARE</summary>
<br>

`Http Request` → `LoggingMiddleware` → `SessionMiddleware` → `AuthMiddleware` → `CustomMiddleware` → `API Server`
<br>

Ketika Http Request sudah sampai ke API server setelah itu API Server akan mengembalikan atau meresponse resquest dari Http Request.
<br>

`API Server` → `CustomMiddleware` → `AuthMiddleware` → `SessionMiddleware` → `LoggingMiddleware` → `Http Response`

</details>

<details>
<summary>3. ECHO MIDDLEWARE</summary>
<br>

Kita bisa melihat Dokumentasi di website khusus di Echo Middleware [https://echo.labstack.com/middleware/](https://echo.labstack.com/middleware/)

<br>

## Setup Middleware Echo

**Echo#Pre() :**

Mengeksekusi sebelum router memproses ke request

- HTTPSRedirect
- HTTPSWWWRedirect
- WWWRedirect
- AddTrailingSlash
- MethodOverride
- Rewrite

**Echo#User() :**

Mengeksekusi setelah router memproses ke request dan memiliki full akses ke echo.Context API

- BodyLimit
- Logger
- Gzip
- Recover
- BasicAuth
- JWTAuth
- Secure
- CORS
- Static
</details>

<details>
<summary>4. TYPE ECHO MIDDLEWARE</summary>

<br>

```go
e :- echo.New()
e.Pre(middleware.HTTPSRedirect())
```

</details>
