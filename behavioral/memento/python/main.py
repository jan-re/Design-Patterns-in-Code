
class Originator():
    def __init__(self, state: str) -> None:
        self._state = state
        print("Initial state: {0}".format(self._state))

    @property
    def state(self) -> str:
        return self._state
    
    @state.setter
    def state(self, new_state: str) -> None


if __name__ == "__main__":
    pass
