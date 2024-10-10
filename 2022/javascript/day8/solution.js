const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

class Point {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    add(other) {
        return new Point(this.x + other.x, this.y + other.y);
    }

    toString() {
        return `${this.x},${this.y}`;
    }
}

function load(file) {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .split("\r\n"));
}

function isVisible(grid, width, height, pos, dir) {
    let nextPos = pos.add(dir);
    while (
        nextPos.x >= 0 &&
        nextPos.x < width &&
        nextPos.y >= 0 &&
        nextPos.y < height
    ) {
        if (grid[nextPos] >= grid[pos]) {
            return false;
        }
        nextPos = nextPos.add(dir);
    }
    return true;
}

function getScenicScore(grid, width, height, pos, dir) {
    let nextPos = pos.add(dir);
    total = 0;
    while (
        nextPos.x >= 0 &&
        nextPos.x < width &&
        nextPos.y >= 0 &&
        nextPos.y < height
    ) {
        total += 1;
        if (grid[nextPos] >= grid[pos]) {
            break;
        }
        nextPos = nextPos.add(dir);
    }
    return total;
}

function solve(p, part1) {
    const width = p[0].length;
    const height = p.length;

    const directions = [
        new Point(1, 0),
        new Point(-1, 0),
        new Point(0, 1),
        new Point(0, -1),
    ];

    let grid = {};
    for (let y = 0; y < p.length; ++y) {
        for (let x = 0; x < p[0].length; ++x) {
            grid[new Point(x, y)] = Number(p[y][x]);
        }
    }

    if (part1) {
        var total = 2 * p[0].length + 2 * (p.length - 2);
    } else {
        var total = 0;
    }

    for (let y = 1; y < p.length - 1; ++y) {
        for (let x = 1; x < p[0].length - 1; ++x) {
            var dir_visible = {};
            for (dir of directions) {
                if (part1) {
                    dir_visible[dir] = isVisible(
                        grid,
                        width,
                        height,
                        new Point(x, y),
                        dir
                    );
                } else {
                    dir_visible[dir] = getScenicScore(
                        grid,
                        width,
                        height,
                        new Point(x, y),
                        dir
                    );
                }
            }
            if (part1 && Object.values(dir_visible).some((value) => value)) {
                total += 1;
            } else if (!part1) {
                scenicScore = Object.values(dir_visible).reduce(
                    (sum, val) => sum * val,
                    1
                );
                scenicScore > total ? (total = scenicScore) : (total = total);
            }
        }
    }

    return total;
}

function main() {
    const startTime = performance.now();
    const puzzle = load("./input.txt");
    const solPart1 = solve(puzzle, true);
    const solPart2 = solve(puzzle, false);

    console.log("Solution Part 1:", solPart1);
    console.log("Solution Part 2:", solPart2);
    console.log(
        "Solved in " +
            ((performance.now() - startTime) / 1000).toFixed(5) +
            " Sec."
    );
}

main();
