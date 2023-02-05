interface FleetFactory {
    deployFlagship(): Flagship
    deployBoardingParty(): BoardingParty
}

class ImperialFleetFactory implements FleetFactory {
    deployFlagship(): Flagship {
        return new ImperialFlagship()
    }

    deployBoardingParty(): BoardingParty {
        return new ImperialBoardingParty()
    }
}

class OrkFleetFactory implements FleetFactory {
    deployFlagship(): Flagship {
        return new OrkFlagship()
    }

    deployBoardingParty(): BoardingParty {
        return new OrkBoardingParty()
    }
}

interface Flagship {
    fireBroadside(): string
}

class ImperialFlagship implements Flagship {
    fireBroadside(): string {
        return 'Gloriana-class battleship unleashes a massive broadside.'
    }
}

class OrkFlagship implements Flagship {
    fireBroadside(): string {
        return 'Ork flagship fires a bunch of garbage in the vague direction of your target.'
    }
}

interface BoardingParty {
    reportMen(): number
}

class ImperialBoardingParty implements BoardingParty {
    reportMen(): number {
        return 8
    }
}

class OrkBoardingParty implements BoardingParty {
    reportMen(): number {
        return 30
    }
}

class NavalEngagement {
    /**
     * Essential segment.
     * Abstract factory field.
     * Any factory fulfilling the interface can be inserted here and worked with below.
     */
    factory: FleetFactory

    constructor(f: FleetFactory) {
        this.factory = f
    }
    
    resolve(): void {
        const flagship = this.factory.deployFlagship()
        const boardingParty = this.factory.deployBoardingParty()

        console.log(flagship.fireBroadside())
        console.log(`You have ${boardingParty.reportMen()} available boarders.`)
    }
}

function main(faction: string) {
    if (faction === 'Imperial') {
        var fleetFactory = new ImperialFleetFactory()
    } else if (faction === 'Ork') {
        var fleetFactory = new OrkFleetFactory()
    } else {
        console.log('Illegal faction. Exiting.')
        return
    }

    const newEngagement = new NavalEngagement(fleetFactory)

    newEngagement.resolve()
}

main('Ork')