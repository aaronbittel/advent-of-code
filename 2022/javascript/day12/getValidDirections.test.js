const path = require("path");
const { getValidDirections, toId } = require(path.join(
    __dirname,
    "solution.js"
));

describe("filterDirectionsThatAreOutOfGrid", () => {
    const grid = [
        [1, 2, 3],
        [2, 5, 6],
        [7, 8, 9],
    ];

    it("should return {1 0}, {0 1}", () => {
        const curr = { x: 0, y: 0 };

        const result = getValidDirections(grid, curr, new Set());

        expect(result).toEqual([
            { x: 1, y: 0 },
            { x: 0, y: 1 },
        ]);
    });
});

describe("filterDirectionsIfStepTooSteep", () => {
    const grid = [
        [1, 4, 3],
        [2, 5, 6],
        [7, 8, 9],
    ];

    it("should return {0 -1}, {1, 0}", () => {
        const curr = { x: 1, y: 1 };

        const result = getValidDirections(grid, curr, new Set());

        expect(result).toEqual([
            { x: 0, y: -1 },
            { x: 1, y: 0 },
        ]);
    });
});
