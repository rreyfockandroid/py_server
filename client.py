import grpc

import api_pb2 as api
import api_pb2_grpc as api_grpc
from utils.clock import wrapper


@wrapper
def run():
    channel = grpc.insecure_channel('localhost:50051')
    stub = api_grpc.PersonServiceStub(channel)
    try:
        response = stub.GetPerson(api.Person(id=1))
        print(response)
    except grpc.RpcError as e:
        print(e.code())


if __name__ == '__main__':
    run()