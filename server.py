import api_pb2 as api
import api_pb2_grpc as api_grpc

from datasource.person_ds import PersonDS

import grpc
import time
import signal

from grpc import StatusCode
from concurrent import futures

from utils.clock import wrapper

CONST_PORT = 50051

class PersonServiceServicer(api_grpc.PersonServiceServicer):

    person_ds = PersonDS()

    @wrapper
    def GetPerson(self, person_api, context):
        person = self.person_ds.get_person(person_api.id)
        if person:
            return api.PersonDetails(id=person.id, name=person.name, email=person.email)
        context.abort(StatusCode.NOT_FOUND, 'Person not found')
                          

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    api_grpc.add_PersonServiceServicer_to_server(PersonServiceServicer(), server)
    server.add_insecure_port(f'[::]:{CONST_PORT}')
    server.start()

    def exit(a, b):
        print('sigint bug - exit')
        server.stop(0)
    # fix for ctrl+c
    signal.signal(signal.SIGINT, exit)

    server.wait_for_termination()

if __name__ == '__main__':
    print('start server ', CONST_PORT)
    serve()