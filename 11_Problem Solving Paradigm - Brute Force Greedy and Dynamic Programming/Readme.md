# RANGKUMAN PROBLEM SOLVING PARADIGM

<details open>
<summary>1. APA ITU PROBLEM SOLVING PARADIGM</summary>
<br>

**Problem Solving Paradigm** adalah pendekatan yang biasa digunakan untuk memecahkan masalah: _Complete Search (a.k.a. Brute Force)_, _Devide and Conquer_, _Greedy_, dan _Dynamic Programming_. Setiap Masalah harus diselesaikan dengan pendekatan yang dapat diselesaikan.

</details>

<details>
<summary>2. BRUTE FORCE</summary>
<br>

**Complete Search (Brute Force)** adalah metode untuk menyelesaikan masalah dengan menelusuri seluruh ruang pencarian untuk mendapatkan solusi yang dibutuhkan. Brute Force terjadi ketika tidak ada algoritma lain yang tersedia, biasanya mudah untuk ditulis karena sangat mudah.

### Contoh Code dari Complete Search :

```go
func minmax(n []int) (int, int) { //brute force (dicek semua)
	var minim = n[0]
	var maxim = n[0]

	for i := 1; i < len(n); i++ {
		if n[i] < minim {
			minim = n[i]
		}

		if n[i] > maxim {
			maxim = n[i]
		}
	}
	return minim, maxim
}

func main() {
	fmt.Print("Min Max :\n")
	fmt.Println(minmax([]int{11, 12, 13, 6, 1, 20, 15}))
	fmt.Println()
}
```

</details>

<details>
<summary>3. DEVIDE AND CONQUER</summary>
<br>

**Divide and Conquer** adalah paradigma pemecahan masalah di mana sebuah masalah dibuat dengan cara menyederhanakannya menjadi bagian-bagian yang lebih kecil dan kemudian menaklukkan setiap bagian :

- **Devide** : Membagi masalah yang besar cukup kecil untuk yang lebih kecil
  <br>

- **Conquer** : Ketika masalah sudah cukup kecil untuk diselesaikan, langsung diselesaikan
  <br>

- **Combine** : Jika dibutuhkan maka perlu menggabungkan solusi dari masalah - masalah yang lebih kecil mejadi solusi untuk masalah yang besar
</details>

<details>
<summary>4. GREEDY</summary>
<br>

**Greedy** adalah sebuah algoritma dikatakan serakah jika algoritma tersebut membuat pilihan yang optimal secara lokal pada setiap langkah dengan harapan pada akhirnya mencapai solusi yang optimal secara global. dalam beberapa kasus, algoritma serakah berhasil - solusinya shrot dan berjalan secara efisien.

</details>

<details>
<summary>5. DYNAMIC PROGRAMMING</summary>
<br>

**Dyanamic Programming** adalah teknik algoritmik untuk memecahkan masalah optimasi dengan memecahnya menjadi submasalah yang lebih sederhana dan memanfaatkan fakta bahwa solusi optimal untuk keseluruhan masalah tergantung pada solusi optimal untuk submasalahnya.

</details>
