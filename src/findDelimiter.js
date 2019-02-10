function findDelimiter(input) {
    if (!input.match(/^\/\//))
        return [undefined, input]
    const [delimCommand, cleanInput] = input.split(/\n/)
    return [
        delimCommand[2],
        cleanInput
    ]
}

function createStringCalculator() {
    return {
        add(input) {
            if (!input)
                return 0
            const [delim, cleanInput] = findDelimiter(input)
            const regex = new RegExp(`${delim || ','}|\n`, 'g')
            return cleanInput.split(regex)
                .map(n => {
                    const num = parseInt(n, 10)
                    if (num < 0)
                        throw new Error(`negatives not allowed: ${n}`)
                    return num
                })
                .reduce((acc, x) => {
                    if (x > 1000)
                        return acc
                    return acc + x
                }, 0)
        }
    }
}

exports.findDelimiter = findDelimiter
exports.createStringCalculator = createStringCalculator
