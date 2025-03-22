import api_pb2 as api
import api_pb2_grpc as api_grpc

import time
import grpc
from grpc import StatusCode
from grpc.aio import server
import asyncio

CONST_PORT = 50051

persons = {
    1: api.PersonDetails(id=1, name='Jan'),
    2: api.PersonDetails(id=2, name='Anna'),
    3: api.PersonDetails(id=3, name='Krzysztof'),
    4: api.PersonDetails(id=4, name='Karol'),
}

class PersonServiceServicer(api_grpc.PersonServiceServicer):

    async def GetPerson(self, person, context):
        start = time.time()
        await asyncio.sleep(1)
        if person.id in persons:
            await asyncio.sleep(1)
            end = time.time()
            print(f'--Execution time: {end- start}')
            return persons[person.id]
        await context.abort(StatusCode.NOT_FOUND, 'Person not found')

async def serve():
    server_instance = server()
    api_grpc.add_PersonServiceServicer_to_server(PersonServiceServicer(), server_instance)
    server_instance.add_insecure_port(f'[::]:{CONST_PORT}')
    await server_instance.start()
    print(f'Server started on port {CONST_PORT}')
    try:
        await server_instance.wait_for_termination()
    except KeyboardInterrupt:
        await server_instance.stop(0)
    except Exception as e:
        await server_instance.stop(1)

if __name__ == '__main__':
    asyncio.run(serve())