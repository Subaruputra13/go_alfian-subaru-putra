# RANGKUMAN STRING, ADVANCE FUNCTION , POINTER , METHOD, STRUCT DAN INTERFACE

<details open>
<summary>1. STRING</summary>
<br>

Pada pembahasan ini String mempunyai Package yang sudah di sediakan oleh Golang, yang biasa di pakai yaitu :

- **Len String** = mengetahui panjang string yang ditentukan

  ```go
    sentence := "Hello";
    lenSentence := len(sentence)
    fmt.Println(lenSentence)
  ```

  <br>

- **Compare String** = untuk mengetahui apakah kata nya sama atau tidak yang bernilai boolean

  ```go
    str1 := "abc"
    str2 := "abd"
    fmt.Println(str1 == str2)
  ```

  <br>

- **Containts String** = untuk mengetahui apakah kata yang sama pada 2 string

  ```go
     res := strings.Contains(str, substr)
    fmt.Println(res) // true
  ```

  <br>

- **Substring** = untuk mengetahui mengambil bagian dari suatu string

  ```go
    value := "cat;dog"
    // Take substring from index 4 to length of string.
    substring := value[4:len(value)]
    fmt.Println(substring)
  ```

  <br>

- **Replace String** = untuk menghapus bagian dari suatu string

  ```go
    // 5. Replace
    s := "this[things]I would like to remove"
    t := strings.Replace(s, "[", "", -1)
    fmt.Printf("%s\n", t)
  ```

  <br>

- **Insert String** = untuk memasukan sebuah kata ke String

  ```go
    // 6. Insert
    p := "green"
    index := 2
    q := p[:index] + "HI" + p[index:]
    fmt.Println(p, q)
  ```

  </details>

<details>
<summary>2. ADVANCE FUNCTION</summary>
<br>

### Variadic Function

Berfungsi untuk :

- Untuk melewatkan pembuatan Slice untuk meneruskan ke func
- Ketika jumlah parameter di input tidak dapat diketahui
  <br>

```go
package main

import (
  "fmt"
)

func sum(numbers  ...int) int { //variadic

  var total int = 0
  for _, number := range numbers {
    total += number
  }
  return total
}

func main() {
  avg := sum(2, 4, 3, 5)
  fmt.Println(avg)
}
```

<br>

### Anonymous Function == Literal Function

Anonymous Function adalah fungsi yang tidak mengandug nama apapun
<br>

Contoh Code :

```go
package main

import "fmt"

func main() {
  // Anonymous function
  func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }()

  // Assigning an anonymous function to a variable
  value := func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }
  value()

  // Passing arguments in anonymous function
  func(sentence string) {
    fmt.Println(sentence)
  }("GeeksforGeeks")
}
```

<br>

### Closure

**Closure** adalah tipe khusus dari anonymous function yang mereferensikan di luar fungsi itu sendiri. Dalam hal ini kita akan menggunakan variabel yang tidak diteruskan ke fungsi sebagai parameter, melainkan tersedia saat fungsi dideklarasikan.
<br>

Contoh Code :

```go
package main

import "fmt"

func newCounter() func() int {
  count := 0 // Closure
  return func() int {
    count += 1
    return count // Closure
  }
}

func main() {
  counter := newCounter()
  fmt.Println(counter())
  fmt.Println(counter())
}
```

<br>

### Defer Function

**Defer Function** adalah fungsi yang hanya di jalakan setelah fungsi induknya dikembalikan. Pengembalian berganda juga dapat digunakan, mereka dijalankan sebagai tumpukan, satu per satu.
<br>

Contoh Code :

```go
package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("Later")
  }() //Defer Func

  fmt.Println("First") //dijalakan Pertama

	// Setelah itu Defer Func dijalankan di sini
```

</details>

<details>
<summary>3. POINTER</summary>
<br>

**Pointer** adalah variable yang menyimpan alamat memori dari variable lain. Pointer memiliki kekuatan untuk mengubah data yang kita tuju. Memory adalah urutan kotak ditempatkan satu demi satu dalam satu baris
<br>

Contoh Pointer Declaration :

```go
package main

import "fmt"

func main() {
  var name string = "John"
  var nameAddress *string = &name
  fmt.Println("name (value)   :", name)                // John
  fmt.Println("name (address) :", &name)               // 0xc000010050
  fmt.Println("nameAddress (value)   :", *nameAddress) // John
  fmt.Println("nameAddress (address) :", nameAddress)  // 0xc000010050
}
```

<br>

untuk mengubah variable dengan memory yang sama dengen code berikut ini :

```go
package main

import "fmt"

func main() {
  var name string = "John"
  var nameAddress *string = &name
  fmt.Println("name (value)   :", name) // John
  fmt.Println("name (address) :", &name) // 0xc20800a220
  fmt.Println("nameAddress (value)   :", *nameAddress) // John
  fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220

  name = "Doe"

  fmt.Println("")
  fmt.Println("name (value)   :", name) // Doe
  fmt.Println("name (address) :", &name) // 0xc20800a220
  fmt.Println("nameAddress (value)   :", *nameAddress) // Doe
  fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220
}
```

<br>

### 2 Importan Operator di Pointer :

- `*` Operator :

  - Mencetak pointer varibale
  - akses value yang di simpan di address

- `&` Operator :
  - Mengembalikan alamat variabel
  - akses address variable ke pointer

Contoh Code Zero Value Pointer `<nil>` :

```go
package main

import (
  "fmt"
)

func main() {
  number_a := 25
  var number_b *int
  if number_b == nil {
    fmt.Println("number_b is", number_b)
    number_b = &number_a
    fmt.Println("number_b after init : is", *number_b)
  }
}
```

<br>

OUTPUT :

```go
Output :

number_b is <nil>
number_b after init : is 25
```

<br>

Contoh Code Pointer Declaration dengan Built-In New() :

```go
package main
import (
  "fmt"
)

func main() {
  var size = new(int)
  fmt.Printf("Size value is %d \n", *size)
  fmt.Printf("Type is %T \n", size)
  fmt.Printf("Address is %v \n", size)
  *size = 85
  fmt.Println("New size value is", *size)
}
```

<br>

OUTPUT :

```go
Output :

Size value is 0
Type is *int
Address is 0xc00007c008
New size value is 85
```

</details>

<details>
<summary>4. METHOD</summary>
<br>

**Method** adalah fungsi yang melekat pada suatu tipe (bisa berupa struct atau tipe data)
<br>

### Contoh Declarationnya :

```go
func (receiver StructType) MethodName(parameterList) (returnTypes) {
// block statement
}
```

<br>

### Method Vs Function

```go
func (receiver StructType) functionName(input type) returnType {
  // block statement method
}

func functionName(input type) returnType {
  // block statement function
}
```

<br>

### Cara membuat Method yang simple :

```go
package main

import "fmt"

type Employee struct {
    FirstName, LastName string
}

func (e Employee) fullName() string {
    return e.FirstName + " " + e.LastName
}

func main() {
    e := Employee{
        FirstName: "Ross",
        LastName:  "Geller",
    }
    fmt.Println(e.fullName())
}
```

<br>

### Kenapa Method bukan Fungsi ?

- Membantu kita dalam menulis kode OOP di Golang
- Method membantu kita untuk menghindari Conflict penamaan
- Pemanggilan Method jauh lebih mudah dibaca dan di pahami dari pada pemanggilan fungsi
  <br>

### Contoh Penggunaan Struct untuk OOS(Object Oriented Style) :

```go
package main

import "fmt"

type Person struct {
  name string // Both non exported fields.
  age  int
}

func (P Person) GetName() string {
  return P.name + " amazing!"
}

func (P *Person) IncreaseAge() {
  P.age = P.age + 1
}

func main() {
  PersonA := Person{"John", 50}
  fmt.Printf("%v\n", PersonA)
  fmt.Println(PersonA.GetName())

  PersonA.IncreaseAge()
  fmt.Println(PersonA.age)
}
```

<br>

```go
    Output :

    {John 50}
    John amazing!
    51
```

<br>

### Contoh penggunaan Method yang membantu dari nama yang conflict :

```go
package main

import (
  "fmt"
  "math"
)

type Rect struct {
  width  float64
  height float64
}

type Circle struct {
  radius float64
}

func (r Rect) Area() float64 {
  return r.width * r.height
}

func (c Circle) Area() float64 {
  return math.Pi * c.radius * c.radius
}

func main() {
  rect := Rect{5.0, 4.0}
  cir := Circle{5.0}
  fmt.Printf("Area of rectangle rect = %0.2f\n", rect.Area())
  fmt.Printf("Area of circle cir = %0.2f\n", cir.Area())
}
```

<br>

```go
Output :

Area of rectangle rect = 20.00
Area of circle cir = 78.54
```

<br>

### Contoh penggunaan Struct Pointer Receiver :

```go
package main

import "fmt"

type Employee struct {
  name   string
  salary int
}

func (e *Employee) changeName(newName string) {
  (*e).name = newName
}

func main() {
  e := Employee{
    name:   "Ross Geller",
    salary: 1200,
  }

  // e before name change
  fmt.Println("e before name change =", e)
  // create pointer to `e`
  ep := &e
  // change name
  ep.changeName("Monica Geller")
  // e after name change
  fmt.Println("e after name change =", e)
```

<br>

```go
Output :

e before name change = {Ross Geller 1200}
e after name change = {Monica Geller 1200}
```

</details>

<details>
<summary>5. STRUCT</summary>
<br>

- **Struct** adalah kumpulan definisi variable dan Function, yang dibungkus sebagai tipe data baru dengan nama tertentu.
  <br>

- **Struct** adalah tipe yang ditentukan pengguna yang berisi kumpulan variable atau fungsi(method) bernama.
  <br>

- **Declaration Struct** :

```go
type struct_variable_name struct {
   field <data_type>
   field <data_type>
   ...
   field <data_type>
}
```

<br>

- Initiliasasi dan Akses field :

```go
package main

import "fmt"

type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func main() {

  var Person0 Person
  Person0.FirstName = "Muchson"
  Person0.LastName = "Ibi"
  Person0.Age = 27
  fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)

  // long declaration with assigned value
  var Person1 = Person{"Rizky", "Kurniawan", 26}
  fmt.Println(Person1)

  // long declaration with assigned value each name fields
  var Person2 = Person{
    FirstName: "Iswanul",
    LastName:  "Umam",
    Age:       25,
  }
  fmt.Println(Person2)

  // sort declaration
  Person3 := Person{"Pranadya", "Bagus", 23}
  fmt.Println(Person3)

  // short declaration with new keyword
  Person4 := new(Person)
  Person4.FirstName = "Muhammad"
  Person4.LastName = "Ismail"
  Person4.Age = 30
  fmt.Println(*Person4)

}
```

</details>

<details>
<summary>6. INTERFACE</summary>
<br>

**Interface** adalah kumpulan method yang dapat diimplentasikan objek, Karena Interface mendefinisikan perilaku objek.
<br>

### Contoh Declaration Interface

```go
type interface_name interface {
   method_name1 <return_type>
   method_name2 <return_type>
   method_name3 <return_type>
   ...
   method_namen <return_type>
}
```

<br>

### Contoh Implentasi Interface :

```go
package main

import "fmt"

type calculate interface {
  large() int
}

type square struct {
  side int
}

func (s square) large() int {
  return s.side * s.side
}

func main() {
  var dimResult calculate
  dimResult = square{10}
  fmt.Println("large square :", dimResult.large())
}
```

<br>

```go
Output :

large square : 100
```

<br>

### Contoh Dari Interface Kosong :

```go
package main

import "fmt"

func describe(i interface{}) {
  fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
  var i interface{}
  describe(i)

  i = 42
  describe(i)

  i = "hello"
  describe(i)
}
```

<br>

```go
Output :

(<nil>, <nil>)
(42, int)
(hello, string)
```

<br>

### Contoh Code Type Assertion Interface :

```go
package main

import "fmt"
import "strings"

func main() {
  var secret interface{}

  secret = 2
  var number = secret.(int) * 10
  fmt.Println(secret, "multiplied by 10 is :", number)

  secret = []string{"apple", "manggo", "banana"}
  var gruits = strings.Join(secret.([]string), ", ")
  fmt.Println(gruits, "is my favorite fruits")
}
```

<br>

```go
Output :

2 multiplied by 10 is : 20
apple, manggo, banana is my favorite fruits
```

<br>

### Contoh Code Type Switch Interface :

```go
package main

import (
  "fmt"
  "strings"
)

func explain(i interface{}) {
  switch i.(type) {
  case string:
    fmt.Println("i stored string ", strings.ToUpper(i.(string)))
  case int:
    fmt.Println("i stored int", i)
  default:
    fmt.Println("i stored something else", i)
  }
}

func main() {
  explain("Hello World")
  explain(52)
  explain(true)
}
```

<br>

```go
Output :

i stored string  HELLO WORLD
i stored int 52
i stored something else true
```

<br>
</details>
