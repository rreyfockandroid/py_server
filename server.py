import api_pb2 as api
import api_pb2_grpc as api_grpc

import grpc
import time

from grpc import StatusCode
from concurrent import futures

from utils.clock import wrapper

CONST_PORT = 50051

persons = {
    1: api.PersonDetails(id=1, name='Jan'),
    2: api.PersonDetails(id=2, name='Anna'),
    3: api.PersonDetails(id=3, name='Krzysztof'),
    4: api.PersonDetails(id=4, name='Karol'),
}

class PersonServiceServicer(api_grpc.PersonServiceServicer):

    @wrapper
    def GetPerson(self, person, context):
        # time.sleep(1)
        if person.id in persons:
            # time.sleep(1)
            return persons[person.id]
        context.abort(StatusCode.NOT_FOUND, 'Person not found')
                          

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    api_grpc.add_PersonServiceServicer_to_server(PersonServiceServicer(), server)
    server.add_insecure_port(f'[::]:{CONST_PORT}')
    server.start()
    server.wait_for_termination()

if __name__ == '__main__':
    print('start server ', CONST_PORT)
    serve()