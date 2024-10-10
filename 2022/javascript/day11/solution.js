const fs = require("fs");
const path = require("path");
const { performance } = require("perf_hooks");

function load(file) {
    return (input = fs
        .readFileSync(path.join(__dirname, file), "utf8")
        .toString()
        .split("\r\n\r\n")).map((line) => line.split("\r\n"));
}

class Monkey {
    constructor(itemList, operation, test, monkeyTestTrue, monkeyTestFalse) {
        this.itemList = itemList;
        this.operation = operation;
        this.test = test;
        this.monkeyTestTrue = monkeyTestTrue;
        this.monkeyTestFalse = monkeyTestFalse;
        this.activity = 0;
    }

    operate(item) {
        if (this.operation === "* old") {
            return item ** 2;
        }
        return eval(`${item} ${this.operation}`);
    }

    testing(item) {
        return eval(`${item} % ${this.test} === 0`);
    }

    action(part1, modulo) {
        for (let item of this.itemList) {
            this.activity++;
            if (part1) {
                item = Math.floor(this.operate(item) / 3);
            } else {
                item = this.operate(Math.floor(item % modulo));
            }
            if (this.testing(item)) {
                this.monkeyTestTrue.itemList.push(item);
            } else {
                this.monkeyTestFalse.itemList.push(item);
            }
        }
        this.itemList = new Array();
    }
}

function solve(p, part1) {
    monkeyList = new Array();
    for (mon of p) {
        startingItems = mon[1].split(": ")[1].split(", ").map(Number);
        operation = mon[2].split("old ")[1];
        test = Number(mon[3].split("by ")[1]);
        monkeyList.push(new Monkey(startingItems, operation, test));
    }

    let index = 0;
    for (mon of p) {
        monkeyTestTrueIndex = Number(mon[4][mon[4].length - 1]);
        monkeyTestFalseIndex = Number(mon[5][mon[5].length - 1]);
        monkeyList[index].monkeyTestTrue = monkeyList[monkeyTestTrueIndex];
        monkeyList[index].monkeyTestFalse = monkeyList[monkeyTestFalseIndex];
        index++;
    }

    times = part1 ? 20 : 10000;
    modulo = monkeyList
        .map((monkey) => monkey.test)
        .reduce((mul, val) => val * mul, 1);
    for (let i = 0; i < times; ++i) {
        for (const monkey of monkeyList) {
            monkey.action(part1, modulo);
        }
    }

    return monkeyList
        .map((monkey) => monkey.activity)
        .sort((a, b) => b - a)
        .slice(0, 2)
        .reduce((mul, val) => val * mul, 1);
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
