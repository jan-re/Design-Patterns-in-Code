from __future__ import annotations
from random import sample
from string import ascii_letters, digits
from datetime import datetime
from abc import abstractmethod, ABC


def generate_random_string(length: int = 10) -> str:
    return "".join(sample((ascii_letters + digits), length))


class Originator():

    def __init__(self, state: str) -> None:
        self._state = state
        print("Initial state: {0}".format(self._state))

    @property
    def state(self) -> str:
        return self._state

    @state.setter
    def state(self, new_state: str) -> None:
        self._state = new_state

    # Business logic abstraction that changes the Originator's state.
    def action(self) -> None:
        print("Originator is processing an action that will change its state.")
        self.state = generate_random_string(30)
        print("Originator state is now: {0}".format(self.state))

    def save(self) -> Memento:
        return Memento(self.state)

    def restore(self, memento: AbstractMemento) -> None:
        self.state = memento.state
        print("Originator state has been restored to: {0}".format(self.state))


class AbstractMemento(ABC):

    @abstractmethod
    def get_name(self) -> str:
        pass

    @property
    @abstractmethod
    def date(self) -> str:
        pass

    @property
    @abstractmethod
    def state(self) -> str:
        pass

class Memento(AbstractMemento):

    def __init__(self, new_state: str) -> None:
        self.state = new_state
        self.date = str(datetime.now())[:19]

    @property
    def date(self) -> str:
        return self._date

    @date.setter
    def date(self, date: str) -> None:
        self._date = date

    @property
    def state(self) -> str:
        return self._state

    @state.setter
    def state(self, new_state: str) -> None:
        self._state = new_state

    def get_name(self) -> str:
        return "{0} / {1}".format(self.date, self.state)


class Caretaker():

    def __init__(self, originator: Originator) -> None:
        self.mementos = []
        self.originator = originator

    @property
    def mementos(self) -> list[AbstractMemento]:
        return self._mementos

    @mementos.setter
    def mementos(self, new_mementos: list[AbstractMemento]) -> None:
        self._mementos = new_mementos

    @property
    def originator(self) -> Originator:
        return self._originator

    @originator.setter
    def originator(self, new_originator: Originator) -> None:
        self._originator = new_originator

    def backup(self) -> None:
        print("\nCaretaker: Saving Originator's state...")
        self.mementos.append(self.originator.save())

    def undo(self) -> None:
        try:
            memento = self.mementos.pop()
            print("Caretaker: Restoring state to: {0}".format(memento.get_name()))
            self.originator.restore(memento)
        except IndexError:
            print("Caretaker: No undo point available. Cannot proceed.")

    def show_history(self) -> None:
        print("Caretaker: Here's the list of mementos:")
        for memento in self.mementos:
            print(memento.get_name())


if __name__ == "__main__":
    originator = Originator("First state ever.")
    caretaker = Caretaker(originator)

    caretaker.backup()
    originator.action()

    caretaker.backup()
    originator.action()

    caretaker.backup()
    originator.action()

    print()
    caretaker.show_history()

    print("\nClient: Now, let's rollback!\n")
    caretaker.undo()

    print("\nClient: Once more!\n")
    caretaker.undo()
