
## Program Push Notifications Berbasis gRPC

 ### 1. Persiapan

Kunjungi bagian [berikut](https://github.com/Ivander112/kalkulator-grpc/blob/main/README.md) pastikan point A dan B sudah dikerjakan. 

 ### 2. Compile FIle protobuf
jalankan command dibawah untuk mengcompile file protobuf dalam bahasa Go

    protoc --go_out=. --go-grpc_out=. push_notifications/rpc-stream/notifications.proto
command ini akan menyimpan file protobuff di folder rpc-stream

jalankan command dibawah untuk mengcompile file protobuf dalam bahasa Go

    python -m grpc_tools.protoc -I push_notifications/rpc-stream/ --python_out=push_notifications/client-stream/ --grpc_python_out=push_notifications/client-stream/ push_notifications/rpc-stream/notifications.proto

command ini akan menyimpan file protobuff di folder client-stream

 ### 3. Jalankan server golang
Masuk ke directory push_notifications melalui terminal

    cd push_notifications

lalu jalankan server dengan command

    go run server-stream/server-stream.go

 ### 4. Jalankan client golang

saat menjalankan client. Harus ada argumen untuk menentukan notifikasi apa yang ingin di streaming. Ada 2  jenis argument yang dapat digunakan. Jika command dijalankan tanpa argument maka program akan error. 

 #### a. Notifikasi nutrisi buah

Yang pertama adalah argument "buah" dengan command seperti berikut

    python3 client-stream/client.py buah
    
command ini akan mengembalikan informasi berupa nama buah beserta nutrisi di dalamnya.

 #### b. Notifikasi daftar universitas di suatu negara
notifikasi kedua adalah argument "universitas" dengan command seperti berikut.

    python3 client-stream/client.py universitas
    
setelah command ini dijalankan maka user akan diminta memasukan nama negara. data yang dikembalikan adalah daftar universitas negraa tersebut. Jika nama negara dikosongkan atau nama negra tersebut tidak ada di dalam API maka program akan error.