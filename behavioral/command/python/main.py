from invoker import Invoker
from command import SimpleCommand, ComplexCommand
from receiver import Receiver

if __name__ == "__main__":
    """
    The client code can parameterize an invoker with any commands.
    """
    simple_command = SimpleCommand("Simple payload")
    complex_command = ComplexCommand(Receiver(), "Send email", "Save report")

    invoker = Invoker(on_start=simple_command, on_finish=complex_command)

    invoker.invoke()