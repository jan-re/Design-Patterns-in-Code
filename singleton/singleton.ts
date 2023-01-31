class Singleton {
    private static connection: Singleton

    private constructor() {}

    public static getConnection(): Singleton {
        if (!Singleton.connection) {
            Singleton.connection = new Singleton();
        }

        return Singleton.connection
    }

    public getConnString(): string {
        return 'psql/123/123:8080'
    }


}


// To be continued