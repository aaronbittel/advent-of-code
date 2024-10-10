const fs = require('fs');
const path = require("path");
const { performance } = require('perf_hooks');


function load(file) {
	return input = fs
		.readFileSync(path.join(__dirname, file), 'utf8')
		.toString()
		.split("\r\n\r\n")
}


function solve(p) {
	elves = p
		.map(
			elf => elf.split("\r\n")
			.map(Number)
			.reduce((sum, val) => sum + val, 0))

	elves.sort((a, b) => b-a)
	return [elves[0], elves.slice(0, 3).reduce((sum, val) => sum + val, 0)]
}


function main() {
	const startTime = performance.now()
	const solution = solve(load("./input.txt"))

	console.log("Solution Part 1:", solution[0])
	console.log("Solution Part 2:", solution[1])
	console.log("Solved in " + ((performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()