import api_pb2 as api
import api_pb2_grpc as api_grpc

import signal
import time
import grpc
from grpc import StatusCode
from grpc.aio import server
import asyncio
from utils.clock import wrapper_async
from datasource.person_ds import PersonDS


CONST_PORT = 50051

class PersonServiceServicer(api_grpc.PersonServiceServicer):

    person_ds = PersonDS()

    @wrapper_async
    async def GetPerson(self, person_api, context):
        person = self.person_ds.get_person(person_api.id)
        if person:
            return api.PersonDetails(id=person.id, name=person.name, email=person.email)
        await context.abort(StatusCode.NOT_FOUND, f'Person {person.id} not found')

async def serve():
    server_instance = server()
    api_grpc.add_PersonServiceServicer_to_server(PersonServiceServicer(), server_instance)
    server_instance.add_insecure_port(f'[::]:{CONST_PORT}')
    await server_instance.start()

    print(f'Server started on port {CONST_PORT}')

    await server_instance.wait_for_termination()

if __name__ == '__main__':
    asyncio.run(serve())