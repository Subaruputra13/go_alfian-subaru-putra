# RANGKUMAN VERSION CONTROL AND BRANCH MANAGEMENT(GIT)

<details open>
<summary>1. VERSIONING</summary>
<br>

**Versioning** adalah bagian yang ada di dalam git, Vesioning adalah cara mengatur versi dari source code pada program dan git adalah Version Control Systm (VCS) yang sudah banyak dipakai para enterprice atau developer untuk mengembangkan softwaree mereka secara berkelompok atau bersama sama. Git dibuat pada tahun 2005 oleh Linus Torvalds yang berasal dari Finlandia. Git akan melacak setiap perubahan yang terjadi didalamnya, maka dari itu banyak developer yang menggunakan software ini untuk kolaborasi atau bekerja sama. Git mempunya fitur bernama commit, commit yang berfungsi sebaga tanda perubahan pada file atau folder tetapi tidak akan terjadi peerubahan pada repository, sederhananya commit adalah catatan perubahan.

</details>

<details>
<summary>2. PENGINSTALLAN GIT DI APPLE DAN WINDOWS</summary>
<br>

**Git Apple :**

1. Download git terbaru melalui Mac di website https://git-scm.com/download/mac
2. Mengikuti tutorialnya untuk menginstall git
3. Buka terminal dan ketik `git --version`
4. kalau muncul versi dari gitu yang di instal berarti penginstallan git pada Mac sudah berhasil
   <br>
   <br>

**Git Windows :**

1. Download git terbaru pada Chorme di website https://git-scm.com/download/win
2. Setelah selesai jalankan installer bernama “Git Setup” lalu ikuti langkahnya dan finish
3. Jalankan command prompt dan ketik `git --version`
4. kalau muncul versi dari gitu yang di instal berarti penginstallan git pada Windows sudah berhasil
</details>

<details>
<summary>3. PERINTAH - PERINTAH DI GIT</summary>
<br>

- `git config` = Perintah untuk mengatur konfigurasi sesuai keinginan. Contoh : buka terminal lalu ketik git `config --global namaUser`
  <br>

- `git init` = Perintaah untuk membuat repository baru, Contohnya Arahkan ke folder project dan bukan terminal ketik `git init`
  <br>

- `git add` = Perintah yang biasa digunakan untuk menambah file ke index/StagingArea. Contoh : lakukan perubahan di repository lalu ketik `git add */namaFile`. `git add .` digunakan jika ingin memasukkan semua perubahan file yang ada di repository tersebut.
  <br>

- `git commit` = Perintah untuk melakukan commit pada perubahan, perubahan yang dilakukan di commit tidak akan langsung masuk ke remote repository. Contoh : ketik `git commit -m "fisrt commit"`
  <br>

- `git status` = Perintah untuk menampilakn daftar file yang berubah. Contoh : ketikkan `git status`
  <br>

- `git checkout` = Perintah untuk membuat branch ataupun berpindah branch(cabang). Contoh : Apabila ingin berpindah dari branch main ke branch development maka gunakan perintah git checkout development
  <br>

- `git pull` = Perintah untuk menggabungkan semua perubahan yang ada di remote repository ke dalam direktori lokal. Contoh: ketik `git pull`
  <br>

- `git merge` = Perintah untuk menggabungkan satu ke branch lain yang aktif. Contoh : `git merge namaBranch`
  <br>

- `git reset` = Perintah untuk mengulang/reset index. Di git reset ini terdapat dua opsi yaitu reset hard ataupun reset soft. Contoh: `git reset --hard`
  <br>

- `git rm` -> Perintah untuk menghapus/remove file dari index dan direkotri. Contoh: `git rm namaFile/namaFolder`
</details>
