import time

def wrapper(func):
    def wrap(*args, **kwargs):
        start = time.time()        
        result = func(*args, **kwargs)
        end = time.time()
        print(f'Execution time: {end - start}')
        return result
    return wrap

def wrapper_async(func):
    async def wrap(*args, **kwargs):
        start = time.time()        
        result = await func(*args, **kwargs)
        end = time.time()
        print(f'Execution time: {end - start}')
        return result
    return wrap