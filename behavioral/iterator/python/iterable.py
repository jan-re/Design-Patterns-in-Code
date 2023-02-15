from collections.abc import Iterable
from typing import Any
from iterator import SequentialIterator

class Collection(Iterable):
    _collection: list[Any]

    def __init__(self, collection: list[Any] = []) -> None:
        self._collection = collection

    def __iter__(self) -> SequentialIterator:
        return SequentialIterator(self._collection)

    def get_reverse_iterator(self) -> SequentialIterator:
        return SequentialIterator(self._collection, reverse=True)

    def add_item(self, item: Any) -> None:
        self._collection.append(item)