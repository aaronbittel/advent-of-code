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
        const halfLength = line.length / 2;
        const firstComp = line.slice(0, halfLength);
        const secondComp = line.slice(halfLength);

        for (const char of firstComp) {
            if (secondComp.includes(char)) {
                if (char === char.toUpperCase()) {
                    return sum + char.charCodeAt() - 65 + 27
                } else {
                    return sum + char.charCodeAt() - 97 + 1
                }
            }
        }
    }, 0)

    
    let part2 = []
    for (let i = 0; i < p.length; i += 3) {
        const first = p[i]
        const second = p[i + 1]
        const third = p[i + 2]
        for (let char of first) {
            if (second.includes(char) && third.includes(char)) {
                if (char === char.toUpperCase()) {
                    part2.push(char.charCodeAt() - 65 + 27)
                    break
                } else {
                    part2.push(char.charCodeAt() - 97 + 1)
                    break
                }
            }
        }
    }

    return [part1, part2.reduce((sum, val) => sum + val, 0)];
}



function main() {
	const startTime = performance.now()
	const [solPart1, solPart2] = solve(load("./input.txt"))

	console.log("Solution Part 1:", solPart1)
    console.log("Solution Part 2:", solPart2)
	console.log("Solved in " + ((performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()