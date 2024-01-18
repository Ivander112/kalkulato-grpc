
## Program kalkulator sederhana berbasis gRPC

  ### A. Persiapan Instalasi Golang

### 1. Instalasi Golang

kunjungi link berikut untuk proses instalasi golang : [install golang](https://go.dev/doc/install)


### 2. Instalasi Proto Compiler

setelah instalasi golang. Perlu juga dilakukan instalasi proto compiler. Proses instalasi dapat dilakukan di link berikut :

[instalasi proto compiler](https://grpc.io/docs/protoc-installation/)

Untuk OS windows jika masih kurang jelas dapat mengikuti langkah berikut: [Instalasi proto compiler windows](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/)


### 3. Clone Repository

jalankan command dibawah untuk clone repository ini:


git clone https://github.com/Ivander112/kalkulator-grpc.git


### 4. Install dependencies

jalankan command dibawah untuk menginstall dependensi yang dibutuhkan:

go get -u

### B. Instalasi Python

### 1. Install Pyenv dan set local python
ikuti langkah langkah dibawah untuk install pyenv
#### Mac/Linux

	[https://github.com/pyenv/pyenv](https://github.com/pyenv/pyenv)

#### Windows

	[https://github.com/pyenv-win/pyenv-win](https://github.com/pyenv-win/pyenv-win)

jika sudah berhasil, install python 3.10.12

	pyenv install 3.10.12
lalu set untuk local pythonnya dengan command ini:

    pyenv local 3.10.12

jalankan command dibawah untuk menjalankan server:

### 2. Setup Virtualenv

install virtualenv

    pip install virtualenv
    
lalu buat virtualenv nya dan aktifkan

    # Buat lingkungan virtual
    virtualenv venv
    
    # Aktifkan lingkungan virtual
    # Windows
    venv\Scripts\activate
    
    # Mac / Linux 
    source venv/bin/activate
    
jika ingin mematikan virtualenv

    # Menonaktifkan lingkungan virtual saat ini
    # Windows
    deactivate
    # Mac / Linux
    exit

### 3. Instalasi Library

install library yang dibutuhkan dengan command ini. pastikan virutalenv sedang aktif:

    pip install -r requirements.txt

### C. Uji coba kalkulator gRPC

### 1. Membuat file proto golang

jalankan command dibawah untuk membuat file protobuffer yang akan digunakan oleh server dan client:

protoc --go_out=. --go-grpc_out=. rpc_function/calculator_rpc/calculator.proto

**Note** : command diatas dibuat meyesuaikan path yang digunakan di repository github. Jika ingin mengubah lokasi file maka ubah "rpc_function/calculator_rpc/calculator.proto" sesuai dengan path dari file proto tersebut

### 2. Membuat file proto python

jalankan command dibawah untuk membuat file protobuffer yang akan digunakan oleh server dan client:

    python -m grpc_tools.protoc -I rpc_function/calculator_rpc/ --python_out=client-python --grpc_python_out=client-python rpc_function/calculator_rpc/calculator.proto

**Note** : command diatas dibuat meyesuaikan path yang digunakan di repository github. Jika ingin mengubah lokasi file maka sesuaikan pathnya

### 3. Jalankan Server Golang

    go run server/server.go

jalankan command dibawah untuk menjalankan client:

### 4. Jalankan Client Golang
Client golang dapat dijalnkan dengan command ini:

    go run client/client.go

Saat menjalankan program client maka user akan diminta memasukan 2 bilangan. Program diatur agar saat pengurangan maka bilangan pertama akan dikurangi bilangan kedua dan saat pembagian maka bilangan pertama akan dibagi oleh bilangan kedua. Program client tidak akan menerima input selain angka

### 5. Jalankan Client python
Client python dapat dijalnkan dengan command ini:

    python3 client-python/client.py

client ini memiliki fungsi yang sama dengan client golang