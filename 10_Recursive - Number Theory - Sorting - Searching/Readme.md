# RANGKUMAN Recursive - Number Theory - Sorting - Searching

<details>
<summary>1. RECURSIVE</summary>
<br>

**Recursive** adalah metode pemecahan masalah komputasi dimana solusinya bergantung pada solusi untuk kasus yang lebih kecil dari masalah yang sama. Recursive memecahkan masalah dengan memanggil fungsi diri nya sendiri dari dalam code kita.

### Recursive Strategi

ada 2 cara untuk memikirkan tentang ketika menggunakan Recursive startegy :

- Base Case :
  Apa Simple case dari problem tersebut ?
  <br>

- Recurrence relation :
  Apa hubungan rekursif dari masalah ini dengan masalah serupa yang lebih kecil ?
  <br>

### Contoh penggunaan recursive dengan Faktorial Matematika :

```go
func faktorial(int n) int {
	if n == 1 {
		return 1
	} else {
		return n * faktorial(n - 1)
	}
}

//5! = 5 x 4 x 3 x 2 x 1 = 120 adalah faktorial dari 5
```

<br>
</details>

<details>
<summary>2. NUMBER THEORY</summary>
<br>

**Number Theory** adalah cabang matematika yang mempelajari bilangan bulat. Ada banyak topik dalam bidang Number Theory, seperti Bilangan Prima, Pembagi Terbesar, Kelipatan Terkecil, Faktorial, Faktor Prima.
<br>

### Contoh Code Numbere Theory dari Faktorial :

```go
func faktorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * faktorial(n - 1 )
	}
}
```

</details>

<details>
<summary>3. SEARCHING</summary>
<br>

**Searching** adalah proses menemukan posisi value tertentu dalam daftar value. Pada Seaching ada 3 cara yaitu :

- Linier Seacrh - O(n)
- Bulitins Search
</details>

<details>
<summary>4. SORTING</summary>
<br>

**Sorting** adalah proses pengurutan data dalam urutan tertentu. Biasanya, kita menyortir berdasarkan nilai elemen. Kita dapat menyortir angka, kata, pasangan. Misalnya, kita dapat menyortir siswa berdasarkan tinggi mereka, dan kita dapat menyortir kota secara alfabetis atau berdasarkan jumlah penduduknya. Urutan yang paling sering digunakan adalah urutan numerik dan alfabetik.
<br>

### Ada 4 Cara Sorting yaitu :

- Selection Sort - O(n^2)
- Counting Sort - O(n + k)
- Merge Sort - O(log n)
- Builtins Sort
</details>
