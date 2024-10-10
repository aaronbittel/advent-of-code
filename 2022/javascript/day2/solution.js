const fs = require('fs');
const path = require("path");
const { performance } = require('perf_hooks');


SHAPES = {
    A: 1,
    B: 2,
    C: 3,
    X: 1,
    Y: 2,
    Z: 3
}


function load(file) {
	return input = fs
		.readFileSync(path.join(__dirname, file), 'utf8')
		.toString()
		.split("\r\n")
        .map((v) => v.split(' '));
}



function solve(p) {
    const part1 = p.map(([enemy, me]) => {
        let [val_enemy, val_me] = [SHAPES[enemy], SHAPES[me]]
        if (val_enemy === val_me) {
            return 3 + val_me
        } else if ((val_enemy - val_me) === -1 || (val_enemy - val_me) === 2){
            return 6 + val_me
        } else {
            return val_me
        }
    });

    const part2 = input.map(([left_shape, right_shape]) => {
        const left = SHAPES[left_shape];
    
        if (right_shape === 'X') {
            // Lose
            let right = left - 1 || 3; // If 0, loop to 3 (paper)
            return right;
        } else if (right_shape === 'Y') {
            // Draw
            return left + 3;
        } else {
            // Win
            let right = (left + 1) % 3 || 3; // If 0, loop to 3 (paper)
            return right + 6;
        }
    });

    return [part1.reduce((sum, val) => sum + val, 0), part2.reduce((sum, val) => sum + val, 0)]
}


function main() {
	const startTime = performance.now()
	const [solPart1, solPart2] = solve(load("./input.txt"))

	console.log("Solution Part 1:", solPart1)
    console.log("Solution Part 2:", solPart2)
	console.log("Solved in " + ((performance.now() - startTime) / 1000).toFixed(5) + " Sec.")
}


main()