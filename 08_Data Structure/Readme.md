# RANGKUMAN DATA STRUCTURE

<details open>
<summary>1. ARRAY</summary>
<br>

**Array** adalah structure data yang berisi sekelompok elemen, dapat berisi satu jenis variable dengan ukurang alokasi tetap. Tipe data yang berbeda dapat ditangani sebagai elemen dalam array seperti Numeric, String, Boolean, Arra. Dimensin array adalah jumlah indeks pada variable array. Array multi demensi(lebih dari satu index, maksimal 7 index). Dalam perhitungan, array sering digunakan untuk operasi matriks.
<br>

### Array Declaration :

Untuk mendeklarasikan array kamu perlu menentukan jumlah elemen yang disimpanny dalam tanda kurung siku `[]` , diikuti dengan jenis elemen yang disimpan array.
<br>

Contoh Array Declaration :

- `Var <variable_name> [<size_of_array>] <tipe_variable>`

```go
import (
  "fmt"
  "reflect"
)

func main(){
    var primes [5]int
    var countries [5]string

    fmt.Println(reflect.ValueOf(primes).Kind())
    fmt.Println(reflect.ValueOf(countries).Kind())
}
```

`Output : array array`
<br>

### Assingment dan Access Array Element

```go
package main

import "fmt"

func main() {
  var countries [2]string

  countries[0] = "India"  // index 0, element pertama
  countries[1] = "Canada" // index 1, element kedua

  fmt.Println(countries[0]) // Mencetakan element pertama
  fmt.Println(countries[1]) // Mencetakan element pertama
}
```

`Output : India Canada`
<br>

### Inisialisasi dengan Array Literal

```go
package main

import "fmt"

func main() {
  odd_numbers := [5]int{1, 3, 5, 7, 9}   // Inisialisasi dengan value
  var even_numbers [5]int = [5]int{0, 2, 4} // Assigment Partial

  fmt.Println(odd_numbers) //Mencetak value array
  fmt.Println(even_numbers) //Mencetak value array
}
```

`Output : array 3`
<br>

### Mendeklarsikan isi Array yang belum ditentukan

```go
    package main

import "fmt"

func main() {
  even_numbers := [5]int{1: 2, 2: 4}// Inisialisasi index Specific{index : value}
  fmt.Println(even_numbers)
}
```

`Output : [ 0 2 4 0 0]`
<br>

### Iterasi Array dengan Perulangan For

```go
primes := [5]int{2, 3, 5}

// Perulangan For cara 1
for index := 0; index < len(primes); index++ {
  fmt.Println(primes[index])
}

// Perulangan For cara 2
for index, element := range primes {
  fmt.Println(index, "=>", element)
}
for _, value := range primes {
  fmt.Println(value)
}

// Perulangan For cara 3
index := 0
for range primes {
  fmt.Println(primes[index])
  index++
}
```

</details>

<details>
<summary>2. SLICE</summary>
<br>

**Slice** adalah struktur data yang berisi sekelompok elemen, dapat berisi satu jenis variable (seperti Array) tetapi memilki ukuran alokasi yang dinamis. Slice sebernarnya bukan Array Dinamic. Slice dideklrasikan seperti array keculia tidak menentukan ukuran apapun dalam tanda kurung `[]`
<br>

### Membuat Slice dari Array

```go
package main

import (
  "fmt"
  "reflect"
)

func main() {
  // Pembuatan Array
  var primes = [5]int{2, 3, 5, 7, 11}

  // Pembuatan slice yang mengambil value dari Array
  var part_primes []int = primes[1:4]

  // menambah data ke slice akan menambah data ke array juga
  fmt.Println(reflect.ValueOf(part_primes).Kind())
  fmt.Println(part_primes)
}
```

`Output : Slice [3 5 7]`
<br>

### Mendeklrasikan Slice

```go
package main

import "fmt"

func main() {
  // long declaration
  var even_numbers []int
  fmt.Printf("elements = %v, len = %d, cap = %d\n", even_numbers, len(even_numbers), cap(even_numbers))

  // long declaration with values
  var odd_numbers = []int{1, 3, 5, 7, 9}
  fmt.Printf("elements = %v, len = %d, cap = %d\n", odd_numbers, len(odd_numbers), cap(odd_numbers))

  // short declaration with values
  numbers := []int{1, 2, 3, 4, 5}
  fmt.Printf("elements = %v, len = %d, cap = %d\n", numbers, len(numbers), cap(numbers))

  // using make function
  var primes = make([]int, 5, 10)
  fmt.Printf("elements = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))
}
```

<br>

### Menggunakan Keyword MAKE()

- `func make([] T, len, cap) []T`
- `Make()` memungkinkan kita untuk membuat irisan ketika mendasari larik tidak ditentukan
  <br>

### Penggunaan APPEND() dan COPY() di Slice

- `APPEND()` : untuk menambahkan kapasitas dari slice atau menggabung 2 slice
- `COPY()` : untuk Menyalin value dari suatu slice ke slice yang baru

```go
package main

import "fmt"

func main() {
  var colors = []string{"red", "green", "yellow"}
  colors = append(colors, "purple")

  copied_colors := make([]string, 10)

  copy(copied_colors, colors) // Menyalin Slice colors ke Slice copied_color
  fmt.Println(copied_colors)
}
```

<br>

### SLICE DENGAN VALUE KOSONG

- Value Kosong di slice adalah Slice NIL

```go
package main

import "fmt"

func main() {
  var primes []int
  fmt.Printf("s = %v, len = %d, cap = %d\n", primes, len(primes), cap(primes))

  if primes == nil {
    fmt.Println("s is nil")
  }
}
```

`Output :`
`s = [], len = 0, cap = 0`
`primes is nil`
<br>

</details>

<details>
<summary>3. MAP</summary>
<br>

**Map** adalah struktur data yang menyimpan data dalam bentuk pasangan kunci dan nilai dimana setiap KEY nya unik

```go
package main

import "fmt"

func main() {
  // Deklarasi Panjang
  var salary = map[string]int{}
  fmt.Println(salary)

  // Dekalrasi Panjang dengan Value
  var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
  fmt.Println(salary_a)

  // Deklarasi Pendek
  salary_b := map[string]int{}
  fmt.Println(salary_b)

  // Penggunaan MAKE
  var salary_c = make(map[string]int)
  salary_c["doe"] = 7000 // assign value
  fmt.Println(salary_c)
}
```

<br>

```go
package main
import "fmt"
func main() {

  // Deklarasikan Panjang dengan value
  var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
  fmt.Println(salary_a, len(salary_a))

  salary_a["nabilah"] = 7000 // assign value
  fmt.Println(salary_a)

  delete(salary_a, "iswanul") // Menghapus value berdasarkan key
  fmt.Println(salary_a)

  value, exist := salary_a["umam"] // Mengecek Key yang sudah ada
  fmt.Println(value, exist)

  for key, value := range salary_a { // Perulangan MAP dengan FOR
    fmt.Println("->", key, value)
  }
}
```

</details>

<details>
<summary>4. FUCNTION</summary>
<br>

**Function** adalah bagian dari kode yang dipanggil dengan nama function itu sendiri. Function adalah cara mudah untuk membagai kode untuk menjadi blok-blok yang berguna. Memudahkan untuk menulis kode yang bersih, rapi, dan modular.
<br>

### Deklarasi Function

- `func <nama_function> () {<Statment>}`
- `func <name_function> () <type_return> { <statements> }`
- `func <name_function> (<parameter>) { <statements> }`
  <br>

### Contoh Penggunaan Function dengan tidak atau menggunakan Parameter

```go
package main

import "fmt"

//Function tidak menggunakan Parameter
func sayHello() {
  fmt.Println("Hello")
}

//Function menggunakan Parameter
func greeting(hour int) {
  if hour < 12 {
    fmt.Println("Selamat Pagi")
  } else if hour < 18 {
    fmt.Println("Selamat Sore")
  } else {
    fmt.Println("Selamat Malam")
  }
}

func main() {
  hour := 15
  greeting(hour)
}
```

<br>

### Function degan Pengembalian value (Single dan Multiple)

```go
package main

import (
  "fmt"
  "math"
)

// singe return value
func calculateSquare(side int) int {
  return side * side
}

// multiple return value
func calculateCircle(diameter float64) (float64, float64) {
  var keliling = math.Pi * math.Pow(diameter/2, 2)
  var luas = math.Pi * diameter
  // return 2 value
  return keliling, luas
}

func main() {
  var side = 5
  wide := calculateSquare(side)
  fmt.Printf("Luas persegi empat: %d \n\n", wide)

  var diameter float64 = 15
  keliling, luas := calculateCircle(diameter)
  fmt.Printf("Luas lingkaran: %.2f \n", keliling)
  fmt.Printf("Keliling lingkaran: %.2f \n", luas)
}
```

<br>

### Function dengan Nama Pengembalian Parameter

```go
package main

import "fmt"

// function having named return parameter
func multiplication(a, b int) (mul int) {
  mul = a * b
  return
}

func main() {
  m := multiplication(5, 5)
  fmt.Println("5 x 5 = ", m)
}
```

<br>

### Type Assignment Otomatis

- `func scale(width, height, scale int) (int, int)`
- `func scale(width int, height int, scale int) (int, int)`

</details>
