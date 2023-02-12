from abc import ABC, abstractmethod
from receiver import Receiver


class Command(ABC):

    @abstractmethod
    def execute(self) -> None:
        pass


class SimpleCommand(Command):
    """
    Simple command whose execution requires no receiver object.
    """
    _payload: str

    def __init__(self, payload: str) -> None:
        self._payload = payload

    def execute(self) -> None:
        print("Executing a simple command with the following payload: {0}".format(self._payload))


class ComplexCommand(Command):
    """
    Complex command who delegates the more intricate business logic to its receiver object.
    """
    _a: str
    _b: str
    _receiver: Receiver

    def __init__(self, receiver: Receiver, a: str, b: str) -> None:

        self._receiver = receiver
        self._a = a
        self._b = b

    def execute(self) -> None:
        print("Complex execution will be done with the Receiver object.")

        print("Executing A...\n")
        self._receiver.executeOnA(self._a)

        print("Executing B...\n")
        self._receiver.executeOnB(self._b)
