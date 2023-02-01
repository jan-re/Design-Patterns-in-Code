abstract class Venue {
    protected abstract intro: string

    /**
     * Essential segment.
     * Factory method defined as abstract and filled in through subclasses.
    */
    protected abstract hireMusician(): Musician

    public switchOnRadio() {
        const radio = this.hireMusician()

        console.log(this.intro)
        radio.playMusic()
    }
}

class RockVenue extends Venue {
    constructor() {
        super()
        this.intro = 'This is the Rock venue broadcasting!'
    }

    protected intro: string
    protected hireMusician(): Musician {
        return new RockMusician()
    }
}

class IndieVenue extends Venue {
    constructor() {
        super()
        this.intro = 'Indie broadcast is starting!'
    }

    protected intro: string
    protected hireMusician(): Musician {
        return new IndieMusician()
    }
}

interface Musician {
    playMusic(): void
}

class RockMusician implements Musician {
    playMusic(): void {
        console.log('...If you want these kinds of dreams it\'s Californication...')
    }
}

class IndieMusician implements Musician {
    playMusic(): void {
        console.log('...The wall is bouncing...')
    }
}

function main(): void {
    const rockClub = new RockVenue()
    const indieClub = new IndieVenue()

    rockClub.switchOnRadio()
    indieClub.switchOnRadio()
}

main()