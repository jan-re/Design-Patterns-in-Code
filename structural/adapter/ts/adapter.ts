class CurrentStandard {
    public request(): string {
        return 'Output format produced by the modern code standard.';
    }
}

class LegacyStandard {
    /**
     * This simulates output that is not understandable by current code.
     * Real example could be, for example, XML output in a system now working with JSON
     */
    public legacyRequest(): string {
        return '.dradnatSycageL eht fo roivaheb laicepS';
    }
}

class LegacyAdapter extends CurrentStandard {
    legacy: LegacyStandard

    constructor(ls: LegacyStandard) {
        super()
        this.legacy = ls
    }

    public request(): string {
        const adaptedOutput = this.legacy.legacyRequest().split('').reverse().join('')

        return `Output translated from legacy code: ${adaptedOutput}`
    }
}

function main() {
    const modern = new CurrentStandard()
    console.log(modern.request())
    console.log('###########################')

    const legacy = new LegacyStandard()
    console.log(`Legacy output not understandable: ${legacy.legacyRequest()}`)
    console.log('###########################')

    const legacyAdapted = new LegacyAdapter(legacy)
    console.log(legacyAdapted.request())
}

main()