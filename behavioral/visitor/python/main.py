from __future__ import annotations
from abc import ABC, abstractmethod


class Visitor(ABC):

    @abstractmethod
    def visit_schema(self) -> None:
        pass

    @abstractmethod
    def visit_image(self) -> None:
        pass


class XML_Visitor(Visitor):

    def visit_schema(self) -> None:
        print("Exporting schema into XML...\n")

    def visit_image(self) -> None:
        print("Exporting image into XML...\n")


class PDF_Visitor(Visitor):

    def visit_schema(self) -> None:
        print("Exporting schema into PDF...\n")

    def visit_image(self) -> None:
        print("Exporting image into PDF...\n")


class Component(ABC):

    @abstractmethod
    def accept(self, visitor: Visitor) -> None:
        pass


class SchemaComponent(Component):

    def accept(self, visitor: Visitor) -> None:
        visitor.visit_schema()


class ImageComponent(Component):

    def accept(self, visitor: Visitor) -> None:
        visitor.visit_image()


def run_application(visitor: Visitor, components: list[Component]) -> None:

    for component in components:
        component.accept(visitor)


if __name__ == "__main__":
    schema = SchemaComponent()
    image = ImageComponent()

    components: list[Component] = [schema, image]

    pdf_visitor = PDF_Visitor()
    run_application(pdf_visitor, components)

    xml_visitor = XML_Visitor()
    run_application(xml_visitor, components)
