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
</details>

<details>
<summary>5. STRUCT</summary>
</details>

<details>
<summary>6. INTERFACE</summary>
</details>
