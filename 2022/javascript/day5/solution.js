const fs = require("fs");
const path = require("path");
const perf_hooks = require("perf_hooks");


function load(file) {
    return input = fs.readFileSync(path.join(__dirname, file), 'utf8')
	.toString()
	.split("\r\n\r\n")
    .map(line => line.split("\r\n"))
}


function solve(p, part1) {
    const [plan, moves] = [p[0], p[1]];
    const crates = [...Array((plan[0].length + 1) / 4)].map(_ => Array(0))
    for (line of plan.slice(0, plan.length - 1)) {
        for (let i = 1; i < line.length; i += 4) {
            // check if letter
            if (line[i].toLowerCase() != line[i].toUpperCase()) {
                crates[(i - 1) / 4].unshift(line[i]);
            }
        }
    }

    for (let move of moves) {
        const [count, from, to] = [...move.matchAll("[0-9]+")].map(match => Number(match[0]))
        if (part1) {
            for (let i = 0; i < count; i++) {
                crates[to - 1].push(crates[from - 1].pop())
            }
        } else {
            const supps_to_move = crates[from - 1].splice(-1 * count, count)
            crates[to - 1].push(...supps_to_move)
        }
    }

    return crates.reduce((res, crate) => res + crate[crate.length - 1], "")
}


function main() {
	const startTime = perf_hooks.performance.now()
    puzzle = load("./input.txt")
	const solPart1 = solve(puzzle, true)
    const solPart2 = solve(puzzle, false)

	console.log("Solution Part 1:", solPart1)
    console.log("Solution Part 2:", solPart2)
	console.log("Solved in " + ((perf_hooks.performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()
