import grpc
import calculator_pb2 as pb
import calculator_pb2_grpc as pb_grpc

# Alamat server
server_addr = 'localhost:50055'

def main():
    # Melakukan koneksi ke server
    with grpc.insecure_channel(server_addr) as channel:
        stub = pb_grpc.CalcServiceStub(channel)
        # Input angka dari user
        while True:   
            try:
                n1 = input("Masukkan bilangan pertama: ")
                angka = float(n1)
                break  # Keluar dari loop jika input adalah angka
            except ValueError:
                print("Input bukan angka. Silakan coba lagi.")

        while True:   
            try:
                n2 = input("Masukkan bilangan kedua: ")
                angka = float(n2)
                break  # Keluar dari loop jika input adalah angka
            except ValueError:
                print("Input bukan angka. Silakan coba lagi.")

        req = pb.CalcRequest(operand1=float(n1), operand2=float(n2))

        # Memanggil fungsi CalcAdd dari server
        add_resp = stub.CalcAdd(req)
        print("Hasil penjumlahan:", add_resp.result)

        # Memanggil fungsi CalcSubtract dari server
        subtract_resp = stub.CalcSubtract(req)
        print("Hasil pengurangan:", subtract_resp.result)

        # Memanggil fungsi CalcDivide dari server
        divide_resp = stub.CalcDivide(req)
        if divide_resp.zero_div:
            print("Hasil pembagian: Pembagi 0. Pembagian tidak bisa dilakukan")
        else:
            print("Hasil pembagian:", divide_resp.result)

        # Memanggil metode CalcMultiply dari server
        multiply_resp = stub.CalcMultiply(req)
        print("Hasil perkalian:", multiply_resp.result)

if __name__ == "__main__":
    main()
