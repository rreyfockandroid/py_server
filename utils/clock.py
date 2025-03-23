import time
import asyncio

def wrapper(func):
    def wrap(*args, **kwargs):
        start = time.time()
        # time.sleep(1)       
        result = func(*args, **kwargs)
        end = time.time()
        print(f'Execution time: {end - start}')
        return result
    return wrap

def wrapper_async(func):
    async def wrap(*args, **kwargs):
        start = time.time()
        # await asyncio.sleep(1)      
        result = await func(*args, **kwargs)
        end = time.time()
        print(f'Execution time: {end - start}')
        return result
    return wrap