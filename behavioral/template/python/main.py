from __future__ import annotations
from abc import ABC, abstractmethod


class TemplateClass(ABC):

    def template_method(self) -> None:
        """
        This is the template method the pattern is named after. The order is set.
        Some work is done by the base class. Subclasses have hooks as an option of adding code in select places.
        """
        self.base_operation_1()
        self.required_subclass_operation_1()
        self.base_operation_2()
        self.hook_1()
        self.required_subclass_operation_2()
        self.base_operation_3()
        self.hook_2()

    def base_operation_1(self) -> None:
        print("Template class does bulk of the work in step 1.")

    def base_operation_2(self) -> None:
        print("In step 2, Template class continues doing bulk of the work. Some methods can be overriden by subclasses.")

    def base_operation_3(self) -> None:
        print("Template class finished bulk of the work in step 3.")

    @abstractmethod
    # Method below needs to be overriden by a subclass.
    def required_subclass_operation_1(self) -> None:
        pass

    @abstractmethod
    # Method below needs to be overriden by a subclass.
    def required_subclass_operation_2(self) -> None:
        pass

    def hook_1(self) -> None:
        """
        Hooks are optional. It allows subclasses to add additional code in key places of the algorithm.
        """
        pass

    def hook_2(self) -> None:
        pass


class ConcreteClass_1(TemplateClass):

    def required_subclass_operation_1(self) -> None:
        print("Required operation NUM 1 implemented by the FIRST concrete class.")

    def required_subclass_operation_2(self) -> None:
        print("Required operation NUM 2 implemented by the FIRST concrete class.")


class ConcreteClass_2(TemplateClass):

    def required_subclass_operation_1(self) -> None:
        print("Required operation NUM 1 implemented by the SECOND concrete class.")

    def required_subclass_operation_2(self) -> None:
        print("Required operation NUM 2 implemented by the SECOND concrete class.")

    def hook_1(self) -> None:
        print("Hook_1 overriden by SECOND concrete class.")


if __name__ == "__main__":
    print("Template method subclassed to ConcreteClass_1 now running:\n")
    implementer_1 = ConcreteClass_1()
    implementer_1.template_method()

    print("\n" + ("#" * 50) + "\n")

    print("Template method subclassed to ConcreteClass_2 now running:\n")
    implementer_2 = ConcreteClass_2()
    implementer_2.template_method()
