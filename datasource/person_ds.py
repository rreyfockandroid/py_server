from dataclasses import dataclass
from utils.random import random_id
import math
from multiprocessing import Process, Queue

def is_prime(n: int, q: Queue) -> bool:
    if n < 2:
        q.put(False)
        return False
    if n == 2:
        q.put(True)
        return True
    if n%2 == 0:
        q.put(False)
        return False
    for i in range(3, math.isqrt(n) + 1, 2):
        if n%i == 0:
            q.put(False)
            return False
    q.put(True)
    return True

@dataclass
class PersonDetails:
    id: int
    name: str = ""
    email: str = ""

class PersonDS():
    __persons = {
        1: PersonDetails(id=1, name='Jan'),
        2: PersonDetails(id=2, name='Anna'),
        3: PersonDetails(id=3, name='Krzysztof'),
        4: PersonDetails(id=4, name='Karol'),
        5: PersonDetails(id=5, name='Katarzyna'),
        6: PersonDetails(id=6, name='Kamil'),
    }

    @random_id
    def get_person(self, id):
        
        q = Queue()
        p = Process(target=is_prime, args=(5000111000222021, q))
        p.start()
        p.join()
        print("PrimeResult: ", q.get())

        # is_prime(5000111000222021) # ostro pali proca
        return self.__persons[id] if id in self.__persons else None
