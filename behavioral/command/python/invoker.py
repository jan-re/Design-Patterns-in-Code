from command import Command


class Invoker():
    """
    The Invoker does not work with any receviers directly, only through Command objects.
    """
    _on_start: Command
    _on_finish: Command

    def __init__(self, on_start: Command, on_finish: Command) -> None:
        self._on_start = on_start
        self._on_finish = on_finish

    def invoke(self) -> None:
        print("Running 'On start' commands...")
        if isinstance(self._on_start, Command):
            self._on_start.execute()

        print("Processing middle part of the application (omitted here).")

        print("Running 'On finish' commands...")
        if isinstance(self._on_finish, Command):
            self._on_finish.execute()
