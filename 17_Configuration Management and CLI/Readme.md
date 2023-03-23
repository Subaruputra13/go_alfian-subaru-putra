### COMMAND LINE INTERFACE (CLI)

<details open>
<summary>1. APA ITU COMMAND LINE</summary>
<br>

**Command Line** adalah antarmuka berbasis teks yang cepat, kuat, dan digunakan pengembang untuk berkomunikasi dengan komputer secara lebih efektif dan efisien untuk menyelesaikan tugas yang lebih luas.

</details>

<details>
<summary>2. KENAPA MENGGUNAKAN COMMAND LINE</summary>
<br>

- Granular control pada OS atau Aplikasi
- Manajemen yang lebih cepat untuk sistem operasi dalam jumlah besar
- Kemampuan untuk menyimpan skrip untuk mengotomatiskan tugas reguler
- Pengetahuan antarmuka untuk membantu pemecahan masalah, seperti masalah koneksi jaringan
  <br>

### Command Line Interface :

- Shell = CLI fot OS’ services :

  - UNIX SHELL
  - Command Prompt(MSDOS)
    <br>

- Other app CLI :
  - Python REPL
  - MySQL CLI client
  - Mongo SHELL
  - redis-cli
  </details>

<details>
<summary>3. PENGENALAN UNIX SHELL</summary>
<br>

**Shell** adalah program yang berfungsi sebagai jembatan antara pengguna dan kernel. **Shell Script** adalah bahasa pemrograman yang dikompilasi berdasarkan perintah-perintah shell.
<br>

Shell Legend :

- `me` → your username
- `linuxbox` →. your hostname
- ` → your current path (home)
- `$` → shell for normal user
- `#` → shell for root user
  <br>

### Normal User Vs Root User

- **Normal User** = Diizinkan untuk memanipulasi hanya direktori /home/username
- **Root User** = Diizinkan untuk memanipulasi / dir (Semua direktori)
- **Normal User + Sudoser** = Diizinkan untuk bertindak sebagai root sementara
</details>

<details>
<summary>4. COMMAND POPULER DI UNIX SHELL</summary>
<br>
Directory :

- `pwd` (mengecek direktory yang sedang digunakan)
- `ls` (mengecek list//isi dari direktori yang sedang digunakan)
- `mkdir` (membuat direktori)
- `cd` (masuk direktori)
- `rm` (menghapus file / direktori )
- `cp` (mengcopy file / direktori)
- `mv` (memindahkan file / direktori)
- `ln` (membuat link pada sebuah file / direktori)
  <br>

Files :

- Membuat file : `touch`
- Melihat isi file : `head, cat, tail, less`
- mengedit isi file : `vim , nano`
- mengizinkan permission : `chown. chmod`
- Different : `diff` (membedakan file)
  <br>

Network :

- `ping`
- `ssh`
- `netstat`
- `nmap`
- `ip addr (ifconfig)`
- `wget`
- `curl`
- `telnet`
- `netcat`
  <br>

Utility :

- `man`
- `env`
- `echo`
- `date`
- `which`
- `watch`
- `sudo`
- `history`
- `grep`
- `locate`
</details>
