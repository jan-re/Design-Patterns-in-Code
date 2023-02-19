from __future__ import annotations
from abc import ABC, abstractmethod


class Document:
    """
    The context which has various states. Initiliazed as a draft.
    """

    def __init__(self) -> None:
        self.state = Draft(self)

    @property
    def state(self) -> State:
        """
        Depending on the state object here, the methods below will produce different results.
        They delegate the actual work to the state object.
        """
        return self._state

    @state.setter
    def state(self, new_state: State) -> None:
        print("Context: New state {0}.".format({type(new_state).__name__}))
        self._state = new_state

    def publish_document(self):
        self.state.publish()

    def render_document(self):
        self.state.render()


class State(ABC):
    """
    Abstract state class from which concrete states will inherit.
    """

    def __init__(self, new_document: Document) -> None:
        self._document = new_document

    @property
    def document(self) -> Document:
        return self._document

    @document.setter
    def document(self, new_document: Document) -> None:
        self._document = new_document

    @abstractmethod
    def publish(self) -> None:
        pass

    @abstractmethod
    def render(self) -> None:
        pass


class Draft(State):

    def publish(self) -> None:
        self.document.state = Published(self.document)
        print("Draft successfully published.")

    def render(self) -> None:
        print("Draft needs to be published before rendering. Aborting.\n")


class Published(State):

    def publish(self) -> None:
        print("Saving changes to the published document.")

    def render(self) -> None:
        print("Document successfully rendered.")


if __name__ == "__main__":
    doc = Document()

    print("Attempting to render draft:")
    doc.render_document()

    doc.publish_document()
    print("Attempting to render published doc:")
    doc.render_document()

    doc.publish_document()
