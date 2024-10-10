const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

function load(file) {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .split("\r\n"));
}

function solve(p, part1) {
    const sizes = {};
    const stack = [];

    p.forEach((c) => {
        if (c.startsWith("$ ls") || c.startsWith("dir")) {
            return;
        }

        if (c.startsWith("$ cd")) {
            const dest = c.split(" ")[2];
            if (dest === "..") {
                stack.pop();
            } else {
                const path =
                    stack.length > 0
                        ? `${stack[stack.length - 1]}_${dest}`
                        : dest;
                stack.push(path);
            }
        } else {
            const [size, _] = c.split(" ");
            stack.forEach((path) => {
                sizes[path] = (sizes[path] || 0) + Number(size);
            });
        }
    });

    if (part1) {
        let total = 0;
        for (const [_, size] of Object.entries(sizes)) {
            total += size <= 100000 ? size : 0;
        }

        return total;
    } else {
        const total_disc_space = 70000000;
        const needed_unused_space = 30000000;
        const available_space = total_disc_space - needed_unused_space;
        const directory_space = sizes["/"];
        const to_remove_space = directory_space - available_space;

        let s = directory_space;
        for (const [_, value] of Object.entries(sizes)) {
            if (value > to_remove_space && value < s) {
                s = value;
            }
        }
        return s;
    }
}

function main() {
    const startTime = performance.now();
    const puzzle = load("./input.txt");
    const solPart1 = solve(puzzle, true);
    const solPart2 = solve(puzzle, false);

    console.log(`Solution Part 1: ${solPart1}`);
    console.log(`Solution Part 2: ${solPart2}`);
    console.log(
        `Solved in ${((performance.now() - startTime) / 1000).toFixed(5)} Sec.`
    );
}

main();
