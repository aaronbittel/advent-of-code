mod days;

use crate::days::day4;
use std::fs;

fn main() {
    let input = fs::read_to_string("src/days/day4/input.txt")
        .expect("should have been able to read the file");
    let sol_part1 = day4::part1(&input);
    let sol_part2 = day4::part2(&input);
    println!("Part 1: {sol_part1}");
    println!("Part 2: {sol_part2}");
}
