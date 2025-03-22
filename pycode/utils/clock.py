import time

def wrapper(func):
    def wrap(*args, **kwargs):
        start = time.time()        
        result = func(*args, **kwargs)
        end = time.time()
        print(f'Execution time: {end - start}')
        return result
    return wrap