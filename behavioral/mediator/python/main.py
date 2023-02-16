from __future__ import annotations
from abc import ABC


class Mediator(ABC):

    def notify(self, sender: object, event: str) -> None:
        pass


class ConcreteMediator(Mediator):
    def __init__(self, component1: Component1, component2: Component2) -> None:
        self._component1 = component1
        self._component2 = component2

        self._component1.mediator = self
        self._component2.mediator = self

    def notify(self, sender: object, event: str) -> None:
        if event == "A":
            print("Mediator reacts on A and triggers following operations:")
            self._component2.react()
        elif event == "B":
            print("Mediator reacts on B and triggers following operations:")
            self._component1.react()


class BaseComponent:
    def __init__(self) -> None:
        pass

    @property
    def mediator(self) -> Mediator:
        return self._mediator

    @mediator.setter
    def mediator(self, mediator: Mediator) -> None:
        self._mediator = mediator

    def react(event: str):
        pass


class Component1(BaseComponent):
    def do_a(self) -> None:
        print("Component 1 does A.")
        self.mediator.notify(self, "A")

    def react(event: str):
        print("Component 1 triggers a reaction as a result of event B.")


class Component2(BaseComponent):
    def do_b(self) -> None:
        print("Component 2 does B.")
        self.mediator.notify(self, "B")

    def react(event: str):
        print("Component 2 triggers a reaction as a result of event A.")


if __name__ == "__main__":

    """
    Components c1 and c2 do not communicate together. They only notify the mediator.
    The mediator communicates with c2 as a result of notifications sent by c1 and vice versa.
    """

    c1 = Component1()
    c2 = Component2()

    mediator = ConcreteMediator(c1, c2)

    c1.do_a()
    c2.do_b()
