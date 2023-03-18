# RANGKUMAN Join – Union – Agregasi – Subquery – Function

<details open>
<summary>1. INTRODUCTION DATABASE</summary>
<br>

**Database** merupakan kumpulan data yang tersimpan dalam suatu sistem komputer dan dapat diakses serta dikelola menggunakan perangkat lunak tertentu. Database digunakan untuk menyimpan informasi dalam bentuk tabel dan memungkinkan pengguna untuk melakukan manipulasi data seperti mengambil, memperbarui, atau menghapus data tersebut.

</details>

<details>
<summary>2. RELATION DATABASE</summary>
<br>

Relation Database dibagi menjadi 3 :

- **One to One** : Relasi database satu ke satu adalah hubungan antara dua tabel di mana setiap baris dalam tabel satu memiliki hanya satu pasangan dalam tabel lain, dan setiap baris dalam tabel lain juga hanya memiliki satu pasangan dalam tabel satu. Contoh relasi satu ke satu adalah hubungan antara nomor identitas karyawan dan nomor identitas kartu akses mereka di suatu perusahaan. Setiap karyawan hanya memiliki satu nomor identitas kartu akses, dan setiap nomor identitas kartu akses hanya terkait dengan satu karyawan.
  <br>

- **One to Many** : Relasi database satu ke banyak (one to many) adalah hubungan antara dua tabel di mana setiap baris dalam tabel satu dapat memiliki banyak pasangan dalam tabel lain, tetapi setiap baris dalam tabel lain hanya memiliki satu pasangan dalam tabel satu. Contoh relasi satu ke banyak adalah hubungan antara departemen dan karyawan di suatu perusahaan. Setiap departemen dapat memiliki banyak karyawan, tetapi setiap karyawan hanya terkait dengan satu departemen.
  <br>

- **Many to Many** : Relasi database banyak ke banyak adalah hubungan antara dua tabel di mana setiap baris dalam tabel satu dapat memiliki banyak pasangan dalam tabel lain, dan setiap baris dalam tabel lain juga dapat memiliki banyak pasangan dalam tabel satu. Contoh relasi banyak ke banyak adalah hubungan antara produk dan pesanan di toko online. Setiap pesanan dapat memiliki banyak produk, dan setiap produk dapat terkait dengan banyak pesanan.
</details>

<details>
<summary>3. SQL STATMENT</summary>
<br>

Jenis Perintah Di SQL :

- DML
- DDL
- DCL
  <br>

### Data Defenition Language(DDL)

- `Crete database <database_name>`
- `Use <database_name>`
- `Create table …`
- `Drop table …`
- `Rename table …`
- `ALTER table …`
  <br>

### Data Manipulation Language(DML)

- `Insert into <table_name> values(values1,values2)`
- `Select * from <table_name>`
- `Update from <table_name>`
- `Delete from <table_name>`
  <br>

### Data Control Language(DCL)

- Like / Between
- And / Or
- Order By
- Limit
- Join
- Union
- Agregate (Min, Max, Sum, Avg, Count, Having)
- Subquery
- Function
</details>
