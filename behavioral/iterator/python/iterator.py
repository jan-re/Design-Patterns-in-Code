from collections.abc import Iterator
from typing import Any
from iterable import Collection


class SequentialIterator(Iterator):
    _position: int
    _reverse: bool
    _collection: Collection

    def __init__(self, collection, reverse=False) -> None:
        self._reverse = reverse
        self._collection = collection
        self._position = -1 if reverse else 1

    def __next__(self) -> Any:
        try:
            value = self._collection[self._position]
            self._position += -1 if self._reverse else 1
            
        except IndexError:
            raise StopIteration()

        return value