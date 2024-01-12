## Program kalkulator sederhana berbasis gRPC

### 1.  Instalasi Golang
kunjungi link berikut untuk proses instalasi golang : [install golang](https://go.dev/doc/install)

### 2.  Instalasi Proto Compiler
setelah instalasi golang. Perlu juga dilakukan instalasi proto compiler. Proses instalasi dapat dilakukan di link berikut : 
[instalasi proto compiler](https://grpc.io/docs/protoc-installation/)
Untuk OS windows jika masih kurang jelas dapat mengikuti langkah berikut: [Instalisasi proto compiler windows](https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/) 

### 3. Clone Repository
jalankan command dibawah untuk clone repository ini:

    git clone https://github.com/Ivander112/kalkulator-grpc.git

### 4. Install dependencies
jalankan command dibawah untuk menginstall dependensi yang dibutuhkan:

    go mod download
### 5. Membuat file proto
jalankan command dibawah untuk membuat file protobuffer yang akan digunakan oleh server dan client:

    protoc --go_out=. --go-grpc_out=. rpc_function/calculator_rpc/calculator.proto

**Note** : command diatas dibuat meyesuaikan path yang digunakan di repository github. Jika ingin mengubah lokasi file maka ubah "rpc_function/calculator_rpc/calculator.proto" sesuai dengan path dari file proto tersebut

### 6. Jalankan server dan client
jalankan command dibawah untuk menjalankan server:

    go run server/server.go
jalankan command dibawah untuk menjalankan client:

    go run client/client.go
    
  Saat menjalankan program client maka user akan diminta memasukan 2 bilangan. Program diatur agar saat pengurangan maka bilangan pertama akan dikurangi bilangan kedua dan saat pembagian maka bilangan pertama akan dibagi oleh bilangan kedua. Program client tidak akan menerima input selain angka