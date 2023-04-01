# RANGKUMAN ORM AND CODE STRUCTURE(MVC)

<br>

<details open>
<summary>1. APA ITU OBJECT RELATIONAL MAPPING(GORM)</summary>
<br>

**Object-relation mapping(ORM)** di ilmu komputer adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek.

### DATABASE MIGRATION

<br>

Cara untuk memperbarui versi database agar sesuai dengan perubahan versi aplikasi.Perubahan dapat berupa peningkatan ke versi terbaru atau pengembalian ke versi sebelumnya.
<br>

### KENAPA DB MIGRATION

<br>

Kesederhanaan pembaruan basis data, Kesederhanaan pengembalian basis data Melacak perubahan pada struktur basis data Riwayat struktur basis data ditulis pada kode Selalu kompatibel dengan perubahan versi aplikasi
<br>

### CARA INSTALL GORM

<br>

`go get -u gorm.io/gorm` -> untuk install gorm
`go get -u gorm.io/driver/mysql` -> untuk install driver mysql

</details>

<details>
<summary>2. APA ITU MVC</summary>
<br>

**MVC** adalah kependekan dari Model, View, dan Controller. MVC adalah cara yang populer untuk mengatur kode Anda. Ide besar di balik MVC adalah bahwa setiap bagian dari kode Anda memiliki tujuan, dan tujuan tersebut berbeda-beda.

</details>

<details>
<summary>3. PROJECT STRUCTUR</summary>

### KENAPA MEMBUTUHKAN STRUKTUR :

<br>

- Untuk mencapai aplikasi modular.
- Menerapkan pemisahan masalah.
- Mengurangi konflik pada versi.
</details>
