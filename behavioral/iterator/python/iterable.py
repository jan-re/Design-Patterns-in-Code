from collections.abc import Iterable
from typing import Any

class Collection(Iterable):
    _collection: list[Any]
