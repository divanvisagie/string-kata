const { findDelimiter, createStringCalculator } = require("../src/findDelimiter")

const chai = require('chai')
const expect = chai.expect

function test({ execute, expected, parameters }) {
    describe(`given parameters "${parameters}"`, () => {
        it(`should return ${expected}`, () => {
            const actual = execute(...parameters)
            expect(actual).to.deep.equal(expected)
        })
    })
}

describe('findDelimiter()', () => {
    test({
        parameters: ["1,2,3"],
        execute: findDelimiter,
        expected: [undefined, "1,2,3"]
    })

    test({
        parameters: ["//;\n1;2"],
        execute: findDelimiter,
        expected: [';', '1;2']
    })
})

describe('String calculator', () => {

    describe('add()', () => {
        test({
            parameters: [""],
            execute: createStringCalculator().add,
            expected: 0
        })

        test({
            parameters: ["1"],
            execute: createStringCalculator().add,
            expected: 1
        })

        test({
            parameters: ["1,2"],
            execute: createStringCalculator().add,
            expected: 3
        })

        test({
            parameters: ["1,2,3"],
            execute: createStringCalculator().add,
            expected: 6
        })

        test({
            parameters: ["1\n2,3"],
            execute: createStringCalculator().add,
            expected: 6
        })

        test({
            parameters: ["//;\n1;2"],
            execute: createStringCalculator().add,
            expected: 3
        })

        test({
            parameters: ["2,1001"],
            execute: createStringCalculator().add,
            expected: 2
        })

        describe('given parameters "1,4,-1"', () => {
            it('should throw an exception', () => {
                const {add} = createStringCalculator()
                expect(add.bind(add,"1,4,-1")).to.throw(/^negatives not allowed: -1$/)
            })
        })
    })
})