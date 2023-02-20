from __future__ import annotations
from abc import ABC, abstractmethod


class Context():

    def __init__(self, strategy: Strategy) -> None:
        self._strategy = strategy

    @property
    def strategy(self) -> Strategy:
        return self._strategy

    @strategy.setter
    def strategy(self, new_strategy: Strategy) -> None:
        self._strategy = new_strategy

    def business_logic(self) -> None:
        print("Context: Sorting data using the strategy (not sure how it'll do it)")
        result = self._strategy.sort_list(["d", "b", "c", "e", "a"])
        print(",".join(result))


class Strategy(ABC):

    @abstractmethod
    def sort_list(self, data: list[str]) -> list[str]:
        pass


class OrderedStrategy(Strategy):

    def sort_list(self, data: list[str]) -> list[str]:
        return sorted(data)


class ReversedStrategy(Strategy):

    def sort_list(self, data: list[str]) -> list[str]:
        return sorted(data, reverse=True)


if __name__ == "__main__":
    # Initialize a Context with the ordered strategy.
    context = Context(OrderedStrategy())

    # List will be sorted in ascending order.
    context.business_logic()

    # Replace strategy.
    context.strategy = ReversedStrategy()

    # List will be sorted in descending order.
    context.business_logic()
