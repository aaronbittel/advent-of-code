const fs = require('fs');
const path = require("path");
const { performance } = require('perf_hooks');


function load(file) {
	return input = fs
		.readFileSync(path.join(__dirname, file), 'utf8')
		.toString()
		.split("\n")
}


function find_first_unique_values(str, count) {
    arr = Array.from(str)
    let arrSlice = arr.slice(0, count)
    if (new Set(arrSlice).size === count) {
        return count
    }
    for (let i = count; i < arr.length; ++i) {
        arrSlice.shift()
        arrSlice.push(arr[i])
        if (new Set(arrSlice).size === count) {
            return i + 1
        }
    }
}


function solve(p, count) {
    if (p.length > 1) {
        return solution = p.map(row => find_first_unique_values(row, count))
    } else {
        return solution = find_first_unique_values(p[1], count);
    }
}


function main() {
	const startTime = performance.now()
    const puzzle = load("./input.txt")
	const solPart1 = solve(puzzle, 4)
    const solPart2 = solve(puzzle, 14)

	console.log("Solution Part 1:", solPart1)
    console.log("Solution Part 2:", solPart2)
	console.log("Solved in " + ((performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()