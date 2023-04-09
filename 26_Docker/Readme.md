# RANGKUMAN DOCKER

<details open>
<summary>1. APA ITU DOCKER</summary>
<br>

**Docker** adalah platform open source yang memungkinkan pengembang dan sysadmin untuk membangun, mengemas, dan menjalankan aplikasi di dalam container. Container memungkinkan suatu aplikasi untuk diisolasi dengan lingkungan yang terpisah dan independen, sehingga dapat berjalan di berbagai lingkungan tanpa perlu menginstall dependensi yang berbeda-beda.

<br>

### Continer Vs Virtual Machines

**Container :**

- Pengurangan pada lapisan aplikasi
- Pengurangan pada lapisan aplikasi
  Kontainer menggunakan lebih sedikit ruang daripada VM (gambar kontainer biasanya berukuran puluhan MB)
- Menangani lebih banyak aplikasi dan membutuhkan lebih sedikit VM dan sistem Operationg

**Virtual Machines :**

- Abstraksi hadrware fisik
- Setiap VM menyertakan salinan lengkap sistem operasi
- lambat untuk booting
</details>

<details>
<summary>2. APA ITU CONTAINER</summary>
<br>

Container bukanlah mesin virtual. **Container** adalah sebuah proses dengan isolasi sistem.
<br>

### Docker Basic :

- Image
- Container
- Engine
- Registry
- Control Plane
</details>

<details>
<summary>3. APA YANG KITA BISA LAKUKAN DENGAN DOCKER</summary>
<br>

- Memudahkan pengembangan dan pengujian aplikasi: Docker memungkinkan kita untuk mengemas semua dependensi dan konfigurasi aplikasi ke dalam kontainer yang dapat dijalankan di mana saja. Ini memungkinkan pengembang untuk dengan mudah membuat dan menguji aplikasi di lingkungan yang terisolasi dan konsisten.
- Membuat aplikasi portabel: Kontainer Docker dapat dijalankan di mana saja, termasuk di lingkungan pengembangan, uji coba, dan produksi. Ini membuat aplikasi portabel dan memungkinkan pengembang dan administrator sistem untuk menjalankan aplikasi dengan mudah di berbagai lingkungan.
- Meningkatkan keamanan aplikasi: Kontainer Docker terisolasi satu sama lain dan tidak dapat mengakses sumber daya atau data lain di mesin host. Ini membantu mencegah ancaman keamanan dan memungkinkan administrator sistem untuk lebih mudah mengelola akses dan izin.
- Mengelola aplikasi dengan mudah: Docker menyediakan alat untuk membuat, memperbarui, dan menghapus kontainer dengan mudah. Ini memungkinkan administrator sistem untuk dengan mudah mengelola aplikasi dan infrastruktur dengan efisien.
- Menyederhanakan proses deployment: Docker memungkinkan kita untuk membuat image aplikasi yang dapat diterapkan dengan mudah di berbagai lingkungan. Ini membuat proses deployment menjadi lebih mudah dan dapat membantu mengurangi downtime aplikasi.
</details>

<details>
<summary>4. PERINTAH - PERINTAH DOCKER</summary>
<br>

- `FROM` = Mendapatkan citra dari registri docker
- `RUN` = Menjalankan perintah bash saat membangun kontainer
- `ENV` = Mengatur variabel di dalam kontainer
- `ADD` = Menyalin berkas dengan proses lain
- `COPY` = Menyalin berkas
- `WORKDIR` = Mengatur direktori berkas yang berfungsi
- `ENTRYPOINT` = Menjalankan perintah setelah selesai membangun kontainer
- `CMD` = Menjalankan perintah tetapi dapat ditimpa
</details>
