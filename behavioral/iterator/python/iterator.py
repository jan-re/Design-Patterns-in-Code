from collections.abc import Iterator
from typing import Any

class SequentialIterator(Iterator):
    _position: int
    _collection: list[Any]
    _reverse: bool

    def __init__(self, collection: list[Any], reverse=False) -> None:
        self._collection = collection
        self._position = -1 if reverse else 0

        self._reverse = reverse

    def __next__(self) -> Any:
        try:
            value = self._collection[self._position]
            self._position += -1 if self._reverse else 1
            
        except IndexError:
            raise StopIteration()

        return value
