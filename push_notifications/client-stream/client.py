import grpc
import sys
import notifications_pb2
import notifications_pb2_grpc

def run(option):
    # Memulai koneksi ke server
    channel = grpc.insecure_channel('localhost:50051')
    stub = notifications_pb2_grpc.NotificationsServiceStub(channel)
    # tentukan fungsi yang dipangggil sesuai argumennya
    if option == 'buah':
        response_streams = stub.FruitsNotifications(notifications_pb2.NotificationsRequest())
    elif option == 'universitas':
        country = input("Masukkan nama negaranya: ")  # Input nama negara dari user 
        request = notifications_pb2.NotificationsRequest(Notification_Name=country)
        response_streams = stub.UniversitiesNotifications(request) # Memulai request
    else:
        print("Opsi yang valid adalah 'buah' atau 'universitas'")
        return
    # Mencetak response sesuai notifikasi yang di iginkan
    for response in response_streams:
        if option == 'buah':
            print(f"Fruit: {response.Name}")
            print(f"Calories: {response.Calories}")
            print(f"Carbohydrates: {response.Carbohydrates}")
            print(f"Fat: {response.Fat}")
            print(f"Protein: {response.Protein}")
            print(f"Sugar: {response.Sugar}")
        elif option == 'universitas':
            print(f"University: {response.Name}")
            print(f"Web Pages: {response.Web_Pages}")
        print("-----")

if __name__ == '__main__':
    # menerima argumen dan menunjukan 
    if len(sys.argv) != 2:
        print("masukan argumen'buah' atau 'universitas'")
        sys.exit(1)

    option = sys.argv[1]
    run(option)