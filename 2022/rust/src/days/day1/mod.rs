use std::cmp;

pub fn part1(input: &str) -> i32 {
    let mut max_elf: i32 = 0;
    let elves: Vec<&str> = input.split("\n\n").collect();

    for elf in elves {
        let sum_elf: i32 = elf
            .lines()
            .filter_map(|line| line.parse::<i32>().ok())
            .sum();

        max_elf = cmp::max(max_elf, sum_elf);
    }
    max_elf
}

pub fn part2(input: &str) -> i32 {
    let mut elves = input
        .split("\n\n")
        .map(|elf| elf.lines().map(|c| c.parse::<i32>().unwrap()).sum::<i32>())
        .collect::<Vec<i32>>();

    elves.sort_by(|a, b| b.cmp(a));
    elves.iter().take(3).sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(24_000, part1(&example_input()));
    }

    #[test]
    fn part2_test() {
        assert_eq!(45_000, part2(&example_input()));
    }

    fn example_input() -> String {
        "\
1000
2000
3000

4000

5000
6000

7000
8000
9000

10000"
            .to_string()
    }
}
