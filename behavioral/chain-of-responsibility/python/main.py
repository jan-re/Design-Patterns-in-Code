from __future__ import annotations
from abc import ABC, abstractmethod
from typing import Optional


class Handler(ABC):

    @abstractmethod
    def set_next(self, handler: Handler) -> Handler:
        pass

    @abstractmethod
    def handle(self, request) -> Optional[str]:
        pass


class AbstractHandler(Handler):
    # It is crucial that each handler has a field referencing the next handler.
    _next_handler: Handler = None

    def set_next(self, handler: Handler) -> Handler:
        self._next_handler = handler

        # Returning a handler here is a convenience allowing for easy chaining.
        return handler

    @abstractmethod
    # Concrete handlers will call super() in order to implement this behaviour.
    # Implementing it in a base abstract class allows some code reuse.
    def handle(self, request) -> str:
        if self._next_handler:
            return self._next_handler.handle(request)

        return None


class CatHandler(AbstractHandler):
    def handle(self, request) -> str:
        if request == 'Milk':
            return "Cat drinks the {0}".format(request)

        return super().handle(request)


class DogHandler(AbstractHandler):
    def handle(self, request) -> str:
        if request == 'Bone':
            return "Dog chews the {0}".format(request)

        return super().handle(request)


class MonkeyHandler(AbstractHandler):
    def handle(self, request) -> str:
        if request == 'Banana':
            return "Monkey eats the {0}".format(request)

        return super().handle(request)


def application(handler: Handler) -> None:
    for treat in ['Banana', 'Coffee', 'Milk', 'Bone']:
        eaten = handler.handle(treat)

        if eaten:
            print(eaten)
        else:
            print('Treat {0} wasn\'t eaten.'.format(treat))


if __name__ == "__main__":
    cat = CatHandler()
    dog = DogHandler()
    monkey = MonkeyHandler()

    dog.set_next(cat).set_next(monkey)

    application(dog)
