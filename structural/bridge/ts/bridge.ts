class Abstraction {
    protected implementation: Implementation

    constructor(impl: Implementation) {
        this.implementation = impl
    }

    operation(): string {
        return this.implementation.implementOperation()
    }
}


/**
 * Implementation in a simple example form where only a string is returned.
 * Simplification hides potentially very complex and distinct types of implementations.
 */
interface Implementation {
    implementOperation(): string
}

class ImplementationA implements Implementation {
    implementOperation(): string {
        return "Implementation of type \"A\" that is distinct."
    }
}

class ImplementationB implements Implementation {
    implementOperation(): string {
        return "Implementation of type \"B\" that is peculiar."
    }
}

/**
 * The client app runs blind to the underlying implementation.
 * @param app Abstraction that is constructed with a concrete implementation of the underlying methods.
 */
function runApp(app: Abstraction) {
    console.log(app.operation())
}

function main() {
    const impA = new ImplementationA()
    let abs = new Abstraction(impA)
    runApp(abs)


    // Switch the implementation to type B and run the app again.
    const impB = new ImplementationB()
    abs = new Abstraction(impB)
    runApp(abs)

}

main()