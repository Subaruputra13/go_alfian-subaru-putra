## RANGKUMAN BASIC PROGRAMMING

Pada pertemuan Section 5 mempelajari tentang :

<details open>
<summary>1. PENGENALAN GOLANG</summary>
<br>

- **Golang** adalah bahasa pemrograman yang tujuan umum nya membuat programmer mudah membuat, sederhana dan perangkat lunak yang efisien.
  <br>
- **Golang** adalah bahasa yang bagus untuk program tingkat rendah yang menyediakan layanan ke sistem lain yang disebut System Programming.
  <br>
- **Application Programs** : E-commers, Music Player, Social Media Apps.
  <br>
- **System Programs** : APIs, Game engines, CLI apps.
  <br>
  <br>

### Jalur untuk Mengompiler sebuah program adalah :

Membuat Source Code _===>_ Compile _===>_ Menajalankan Sourcode yang sudah di compile dan akan menghasilkan Output.
<br>
<br>

### Command Go Terminal :

-> `go run` = Menjalankan program tanpa build
<br>

-> `go build` = Mengcompile program
<br>

-> `go install` = Sama seperti go build dan dilanjutkan dengan install process
<br>

-> `go test` = Untuk mengetest dengan suffix_test.go
<br>

-> `go get` = Untuk mendowload Go package

</details>

<details>
<summary>2. PACKAGE "FMT"</summary>
<br>

**Output** :
<br>

- `fmt.Printf()` = Untuk mendeklarasikan format verb
  <br>
- `fmt.Prinln()` = Untuk mencetak baris baru
  <br>
- `fmt.Sprintf()` = Untuk mencetak String dan mendeklarasikan format verb String
  <br>
  <br>

**Scanning** :
<br>

- `fmt.Scanln()` = Untuk membuat inputan
  <br>
  <br>

**Format Verb** : %T, %v, %s, %q, %d

</details>

<details>
<summary>3. VARIABLE, TYPES & ZERO VALUES</summary>
<br>

**Variable** digunakan untuk menyimpan informasi dalam program komputer, mereka menyediakan cara pelabelan data dengan nama deskriptif dan mereka memiliki tipe data (Integer, String, Boolean)
<br>
<br>

**Data Types** :
<br>
-> Boolean = True dan False
<br>
-> Numeric = Interger, Float, Complex
<br>
-> String
<br>
<br>

**Type Declaration** :
<br>

==> `var (variable_name) (type_data)`
<br>

==> `var (variable_name) (type_data) = (value)`
<br>

==> `(variable_name) := (value)`
<br>
<br>

### Zero Value adalah Default Value yang ada di tipe-tipe data

Contohnya :
<br>

- Boolean = false
- <br>
- Float = 0.0
- <br>
- Integer = 0
- <br>
- String = ""
<br>
<br>
</details>

<details>
<summary>4. OPERATOR MATEMATIKA</summary>
<br>

- `+`(Addition) = Operator nambahan
- `-`(Subtraction) = Operator pengurangan
- `/`(Division) = Operator pembagian
- `*`(Multiplication) = Operator perkalian
- `%`(Modulo) = Operator sisa hasil bagi
- `++`(Increment) = Operator untuk menambahan keatas
- `--`(Decremnet) = Operator untuk mengurangi kebawah
  <br>
  <br>

### Operator di Golang :

- Arihmetic = `+`, `-`, `\*`, `/`, `%`, `++`, `--`
  <br>

- Comparison = `==`, `!=`, `>`, `<`, `>=`, `<=`
  <br>

- Logical = `&&`, `||`, `!`
  <br>

- Bitwise = `&`, `|`, `^`, `<<` , `>>`
  <br>

- Assigment = `=`, `+=`, `-=`, `*=`, `/=`, `%=`, `<<=`, `>>=`, `&=`, `^=`, `|=`
  <br>

- Miscellaneous = `&`(array), `*`(pointer)
</details>

<details>
<summary>5. CONTROL STRUCTURE BRANCHING & LOOPING</summary>
<br>

Hanya ada beberapa Struktur Control di Go, untuk bercabang kita menggunakan `IF` dan `SWITCH` , untuk perulangan loop kita menggunakan `FOR`

### Contoh IF Statment :

```go
var umur = 17

if umur >= 5 {
    fmt.Println("Anda Masih Muda")
}

if umurayah := 9; umurayah < umur {
    fmt.Println(umurayah)
}
```

<br>
<br>

### Contoh IF ELSE Statment :

```go
jam := 15

if jam < 12 {
    fmt.Println("Selamat Pagi")
}else if jam < 18 {
    fmt.Println("Selamat Sore")
}else {
    fmt.Println("Selamat malam")
}
```

`Output : Selamat Sore`
<br>
<br>

### Contoh Nested IF Statment :

```go
var v1 int = 400
var v2 int = 700

if v1 == 400 {
    if v2 == 700 {
        fmt.Println("Value of v1 is 400 abd v2 is 700")
    }
}
```

<br>
<br>

### Contoh SWICTH Statment :

```go
var today int = 2

switch today {
    case 1:
        fmt.Println("Today is Monday")
    case 2:
        fmt.Println("Tooday is Tuesday")
    case 3:
        fmt.Println("Invalid Date")
}
```

`Output : Today is Tueday`
<br>
<br>

### Contoh LOOP Statment :

```go
package main

import "fmt"

func main(){
    for i:=0;i < 5;i++{
        fmt.Println(i)
    }
}
```

`Output :`
`0`
`1`
`2`
`3`
`4`
<br>
<br>

</details>
