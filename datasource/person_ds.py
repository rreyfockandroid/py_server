from dataclasses import dataclass
from utils.random import random_id

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
        return self.__persons[id] if id in self.__persons else None

