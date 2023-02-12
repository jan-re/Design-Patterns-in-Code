class Receiver:
    """
    Receiver class designed to process complex business logic for command objects using it.
    """

    def executeOnA(self, a: str) -> None:
        print("Complex payload '{0}' has been processed.".format(a))

    def executeOnB(self, b: str) -> None:
        print("Complex payload '{0}' has been processed with a different method.".format(b))
