/**
 * Essential segment.
 * Builder interface should contain a reset() method and all of the established build steps.
 * Build steps should be common.
 * The eventual products do not have to share an interface.
 */
interface Builder {
    reset(): void
    ConstructHouse(): void
    AddPool(): void
    AddFireplace(): void
    ConstructGarage(): void
    AddGuestBedroom(): void
}

interface House {
    construction: string
    addedFeatures: string[]
}

class LuxuryBuilder implements Builder {
    house: {
        construction: string,
        addedFeatures: string[]
    }

    constructor() {
        this.house = {
            construction: '',
            addedFeatures: []
        }
    }

    reset(): void {
        const house: House = {
            construction: '',
            addedFeatures: []
        }

        this.house = { ...house }
    }

    getProduct(): House {
        return { ... this.house }
    }

    ConstructHouse(): void {
        this.house.construction = 'Built on schedule.'
    }

    AddPool(): void {
        this.house.addedFeatures.push('Pool')
    }

    AddFireplace(): void {
        this.house.addedFeatures.push('Fireplace')
    }
    ConstructGarage(): void {
        this.house.addedFeatures.push('Garage')
    }
    AddGuestBedroom(): void {
        this.house.addedFeatures.push('Guest bedroom')
    }
}

class VirtualBuilder implements Builder {
    compiledTemplate: string[]

    constructor() {
        this.compiledTemplate = []
    }

    reset(): void {
        const compiledTemplate: string[] = []

        this.compiledTemplate = [...compiledTemplate]
    }

    getProduct(): string[] {
        return [...this.compiledTemplate]
    }

    ConstructHouse(): void {
        this.compiledTemplate.push('abstract_house_code')
    }

    AddPool(): void {
        this.compiledTemplate.push('figurative_pool')
    }

    AddFireplace(): void {
        this.compiledTemplate.push('figment_of_fire')
    }
    ConstructGarage(): void {
        this.compiledTemplate.push('imaginary_garage')
    }
    AddGuestBedroom(): void {
        this.compiledTemplate.push('thought_of_guest')
    }
}

class Director {

    createMinimalProduct(b: Builder) {
        b.reset()

        b.ConstructHouse()
        b.AddFireplace()
    }

    createLuxuryProduct(b: Builder) {
        b.reset()

        b.ConstructHouse()
        b.ConstructGarage()
        b.AddFireplace()
        b.AddGuestBedroom()
        b.AddPool()
    }
}

function main() {
    const luxuryBuilder = new LuxuryBuilder()
    const virtualBuilder = new VirtualBuilder()

    const director = new Director()

    director.createLuxuryProduct(luxuryBuilder)
    const luxuryProduct = luxuryBuilder.getProduct()

    director.createMinimalProduct(luxuryBuilder)
    const minimalProduct = luxuryBuilder.getProduct()

    console.log('Luxury physical product:')
    console.log(luxuryProduct)

    console.log('Minimal physical product:')
    console.log(minimalProduct)

    director.createLuxuryProduct(virtualBuilder)
    const virtualLuxury = virtualBuilder.getProduct()

    director.createMinimalProduct(virtualBuilder)
    const virtualMinimal = virtualBuilder.getProduct()

    console.log('Luxury virtual product:')
    console.log(virtualLuxury)

    console.log('Minimal minimal product:')
    console.log(virtualMinimal)
}

main()