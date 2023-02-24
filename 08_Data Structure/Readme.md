_RANGKUMAN DATA STRUCTURE_

<details>
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

`Output : India Canda`
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
</details>

<details>
<summary>3. MAP</summary>
</details>

<details>
<summary>4. FUCNTION</summary>
</details>
