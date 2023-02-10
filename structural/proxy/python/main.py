from abc import ABC, abstractmethod


class Server(ABC):
    @abstractmethod
    def handleRequest(self) -> None:
        pass


class Backend(Server):
    def handleRequest(self) -> None:
        print("Backend: Handling request...")


class Proxy(Server):
    def __init__(self, backend: Backend) -> None:
        self.backend = backend

    def handleRequest(self) -> None:
        print("Proxy logging: Request logged...")
        self.backend.handleRequest()


def application(server: Server):
    server.handleRequest()


if __name__ == "__main__":
    backend = Backend()
    proxy = Proxy(backend)

    print("Handling request with the backend directly.\n")
    application(backend)

    print("Handling request through the proxy.\n")
    application(proxy)
