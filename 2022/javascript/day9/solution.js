const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

class Vector {
    constructor(x, y) {
        this.x = x;
        this.y = y;
    }

    add(other) {
        return new Vector(this.x + other.x, this.y + other.y);
    }

    distance(other) {
        return Math.floor(
            Math.sqrt(
                Math.pow(this.x - other.x, 2) + Math.pow(this.y - other.y, 2)
            )
        );
    }

    follow(other) {
        if (this.x === other.x) {
            return new Vector(0, other.y - this.y > 0 ? 1 : -1);
        } else if (this.y === other.y) {
            return new Vector(other.x - this.x > 0 ? 1 : -1, 0);
        } else {
            const vector = new Vector(other.x - this.x, other.y - this.y);
            return new Vector(vector.x > 0 ? 1 : -1, vector.y > 0 ? 1 : -1);
        }
    }

    equals(other) {
        return this.x === other.x && this.y === other.y;
    }

    toString() {
        return `${this.x},${this.y}`;
    }
}

function addVectorIfNotInSet(set, vector) {
    if (!Array.from(set).some((e) => e.equals(vector))) {
        set.add(vector);
    }
}

function load(file) {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .split("\r\n"));
}

function walk_part1(commands) {
    let unique = new Set([new Vector(0, 0)]);
    let head = (tail = new Vector(0, 0));
    const directions = {
        R: new Vector(1, 0),
        L: new Vector(-1, 0),
        U: new Vector(0, -1),
        D: new Vector(0, 1),
    };

    for (const [cmd, count] of commands) {
        const dir = directions[cmd];
        for (let i = 0; i < count; ++i) {
            head = head.add(dir);
            if (tail.distance(head) <= 1) {
                continue;
            }
            tail = tail.add(tail.follow(head));
            addVectorIfNotInSet(unique, tail);
        }
    }
    return unique;
}

function walk_part2(commands) {
    let unique = new Set([new Vector(0, 0)]);
    let knots = Array.from({ length: 10 }, () => new Vector(0, 0));

    const directions = {
        R: new Vector(1, 0),
        L: new Vector(-1, 0),
        U: new Vector(0, -1),
        D: new Vector(0, 1),
    };

    for (const [cmd, count] of commands) {
        const dir = directions[cmd];
        for (let i = 0; i < count; ++i) {
            knots[0] = knots[0].add(dir);
            for (let i = 1; i < knots.length; ++i) {
                if (knots[i].distance(knots[i - 1]) <= 1) {
                    continue;
                }
                knots[i] = knots[i].add(knots[i].follow(knots[i - 1]));
                if (i == knots.length - 1) {
                    addVectorIfNotInSet(unique, knots[i]);
                }
            }
        }
    }
    return unique;
}

function solve(p, part1) {
    p = p.map((line) => {
        const [cmd, count] = line.split(" ");
        return [cmd, Number(count)];
    });

    if (part1) {
        var unique = walk_part1(p);
    } else {
        var unique = walk_part2(p);
    }
    return unique.size;
}

function main() {
    const startTime = performance.now();
    const puzzle = load("./input.txt");
    const solPart1 = solve(puzzle, true);
    const solPart2 = solve(puzzle, false);

    console.log("Solution Part 1:", solPart1);
    console.log("Solution Part 2:", solPart2);

    const executionTime = ((performance.now() - startTime) / 1000).toFixed(5);
    console.log(`Solved in ${executionTime} Sec.`);
}

main();
