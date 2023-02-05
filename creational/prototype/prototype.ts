class Prototype {
    primitiveExample: string
    objectExample: {
        foo: string
    }
    backreferenceExample: ComponentWithBackReference


    clone(): this {
        const clone = Object.create(this)

        //clone.primitiveExample = this.primitiveExample
        clone.objectExample = Object.create(this.objectExample)
        clone.backreferenceExample = {
            ... this.backreferenceExample,
            prototype: {... this}
        }

        return clone
    }
}

class ComponentWithBackReference {
    public prototype;

    constructor(prototype: Prototype) {
        this.prototype = prototype;
    }
}


function main() {
    const template = new Prototype()

    template.primitiveExample = 'I am a basic string.'
    template.objectExample = {
        foo: 'bar'
    }
    template.backreferenceExample = new ComponentWithBackReference(template)

    const clone = template.clone()

    console.log(`Primitive comparison 'template.primitiveExample === clone.primitiveExample' => ${template.primitiveExample === clone.primitiveExample}`)

    console.log(`Template object: ${template.objectExample['foo']} VS Cloned object: ${clone.objectExample['foo']}`)
    console.log(`Object memory address comparison 'template.objectExample === clone.objectExample' => ${template.objectExample === clone.objectExample}`)

    console.log(`Circular reference comparison 'template.backreferenceExample === clone.backreferenceExample' => ${template.backreferenceExample === clone.backreferenceExample}`)
    console.log(`Circular reference prototype comparison 'template.backreferenceExample === clone.backreferenceExample' => ${template.backreferenceExample.prototype === clone.backreferenceExample.prototype}`)
}

main()


/** Notes
 * With the prototype, key is to establish safe practises for cloning objects within the given programming language.
 * In TS, using the spread operator where appropriate looks to be the simplest way.
 * Risk to keep in mind is the ability to accidentally copy over references to the same memory address as the template.
 */