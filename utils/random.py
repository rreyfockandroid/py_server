
import itertools

def random_id(func):
    count = itertools.count()
    def draw_id(*args, **kwargs):
        i = next(count)
        return func(args[0], i%6+1, **kwargs)
    return draw_id
