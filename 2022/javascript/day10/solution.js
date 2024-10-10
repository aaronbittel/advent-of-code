const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

function load(file) {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .split("\r\n"));
}

function solvePart1(p) {
    let registerValue = 1;
    let cycle = 0;
    const specialCycles = new Map([
        [20, 0],
        [60, 0],
        [100, 0],
        [140, 0],
        [180, 0],
        [220, 0],
    ]);

    let index = 20;
    for (const row of p) {
        if (row.startsWith("addx")) {
            count = Number(row.split(" ")[1]);
            if (cycle + 2 >= index) {
                specialCycles.set(index, registerValue);
                index += 40;
            }
            cycle += 2;
            registerValue += count;
        } else {
            if (cycle + 1 >= index) {
                specialCycles.set(index, registerValue);
            }
            cycle += 1;
        }
        if (index > 220) {
            break;
        }
    }
    let total = 0;
    specialCycles.forEach((value, key) => {
        total += value * key;
    });
    return total;
}

function solvePart2(p) {
    points = new Set();
    let registerValue = 1;
    let crtValue = 0;
    let rowIndex = 0;
    for (const cmd of p) {
        if (cmd.startsWith("addx")) {
            count = Number(cmd.split(" ")[1]);
            for (let i = 0; i < 2; ++i) {
                if (Math.abs(crtValue - registerValue) <= 1) {
                    points.add(crtValue + rowIndex * 40);
                }
                crtValue++;
                if (crtValue === 40) {
                    crtValue = 0;
                    rowIndex++;
                }
            }
            registerValue += count;
        } else {
            if (Math.abs(crtValue - registerValue) <= 1) {
                points.add(crtValue + rowIndex * 40);
            }
            crtValue++;
            if (crtValue === 40) {
                crtValue = 0;
                rowIndex++;
            }
        }
    }

    let part2 = "\n";
    for (let y = 0; y < 6; ++y) {
        for (let x = 0; x < 40; ++x) {
            const check = y * 40 + x;
            if (points.has(check)) {
                part2 += "#";
            } else {
                part2 += ".";
            }
        }
        part2 += "\n";
    }
    return part2;
}

function main() {
    const startTime = performance.now();
    const puzzle = load("./input.txt");
    const solPart1 = solvePart1(puzzle);
    const solPart2 = solvePart2(puzzle);

    console.log("Solution Part 1:", solPart1);
    console.log("Solution Part 2:", solPart2);

    const executionTime = ((performance.now() - startTime) / 1000).toFixed(5);
    console.log(`Solved in ${executionTime} Sec.`);
}

main();
