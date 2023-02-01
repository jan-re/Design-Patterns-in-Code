class Singleton {
    private static connection: Singleton

    private constructor() { }

    public static getConnection(): Singleton {
        if (!Singleton.connection) {
            Singleton.connection = new Singleton();
        }

        return Singleton.connection
    }

    public getConnString(): object {
        return {
            connString: 'psql/123/123:8080'
        }
    }


}

function main() {
    const singleton = Singleton.getConnection()
    const singleton2 = Singleton.getConnection()

    console.log(`Are the two instances the same: ${singleton === singleton2}`)
}

main()