const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

const load = (file) => {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .trim()
        .split("\r\n")
        .map((line) => line.split("")));
};

const getValidDirections = (grid, curr, visited) => {
    return [
        { x: 1, y: 0 },
        { x: -1, y: 0 },
        { x: 0, y: 1 },
        { x: 0, y: -1 },
    ].filter((dir) => {
        const nextPos = { x: curr.x + dir.x, y: curr.y + dir.y };
        return (
            grid[nextPos.y]?.[nextPos.x] !== undefined &&
            !visited.has(toId(nextPos.x, nextPos.y)) &&
            Math.abs(grid[nextPos.y][nextPos.x] - grid[curr.y][curr.x]) <= 1
        );
    });
};

const toId = (x, y) => `${x} ${y}`;

const walk = (grid, start, dest) => {
    let queue = new Array(start);
    let visited = new Set();
    let steps = 0;
    while (queue.length > 0) {
        steps++;
        const curr = queue.shift();
        if (toId(curr.x, curr.y) === toId(dest.x, dest.y)) return steps;
        visited.add(toId(curr.x, curr.y));
        for (const dir of getValidDirections(grid, curr, visited)) {
            queue.push({ x: curr.x + dir.x, y: curr.y + dir.y });
        }
    }
    return steps;
};

const solve = (p) => {
    let start, dest;
    for (let y = 0; y < p.length; ++y) {
        for (let x = 0; x < p[0].length; ++x) {
            const cell = p[y][x];
            if (cell === "S") {
                start = { x, y };
                p[y][x] = "a";
            } else if (cell === "E") {
                dest = { x, y };
                p[y][x] = "z";
            }
            p[y][x] = p[y][x].charCodeAt(0);
        }
    }

    return walk(p, start, dest);
};

const main = () => {
    const startTime = performance.now();
    const puzzle = load("./sample.txt");
    const solPart1 = solve(puzzle);

    console.log("Solution Part 1:", solPart1);

    const executionTime = ((performance.now() - startTime) / 1000).toFixed(5);
    console.log(`Solved in ${executionTime} Sec.`);
};

main();

module.exports = { getValidDirections, toId };
