import grpc
import sys

from concurrent import futures
import time

import notifications_pb2
import notifications_pb2_grpc

def run(option):
    channel = grpc.insecure_channel('localhost:50051')
    stub = notifications_pb2_grpc.NotificationsServiceStub(channel)

    if option == 'buah':
        response_streams = stub.FruitsNotifications(notifications_pb2.NotificationsRequest())
    elif option == 'universitas':
        response_streams = stub.UniversitiesNotifications(notifications_pb2.NotificationsRequest())
    else:
        print("Opsi yang valid adalah 'buah' atau 'universitas'")
        return

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
    if len(sys.argv) != 2:
        print("Usage: python client.py [buah/universitas]")
        sys.exit(1)

    option = sys.argv[1]
    run(option)