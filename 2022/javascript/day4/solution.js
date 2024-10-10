const fs = require('fs');
const path = require("path");
const { performance } = require('perf_hooks');


function load(file) {
	return input = fs
		.readFileSync(path.join(__dirname, file), 'utf8')
		.toString()
		.split("\r\n")
}



function solve(p) {
    const part1 = p.reduce((sum, line) => {
        const [left, right] = line.split(",")
        const [left_from, left_to] = left.split("-").map(Number)
        const [right_from, right_to] = right.split("-").map(Number)

        if (left_from <= right_from && left_to >= right_to || right_from <= left_from && right_to >= left_to) {
            return sum + 1
        } else {
            return sum + 0
        }
    }, 0);

    
    const part2 = p.reduce((sum, line) => {
        const [left, right] = line.split(",")
        const [left_from, left_to] = left.split("-").map(Number)
        const [right_from, right_to] = right.split("-").map(Number)

        if ((left_from < right_from && left_to < right_from) || (left_from > right_from && right_to < left_from)) {
            return sum + 0
        } else {
            return sum + 1
        }

    }, 0);

    
    return [part1, part2]
}



function main() {
	const startTime = performance.now()
	const [solPart1, solPart2] = solve(load("./input.txt"))

	console.log("Solution Part 1:", solPart1)
    console.log("Solution Part 2:", solPart2)
	console.log("Solved in " + ((performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()