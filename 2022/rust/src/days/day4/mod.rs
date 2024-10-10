use std::ops::RangeInclusive;

// dont know what is better: filter or filter_map
pub fn part1(input: &str) -> u32 {
    input
        .lines()
        .filter_map(|line| {
            let (first, second) = line.split_once(',').expect("I expect a comma");
            let (start1, end1) = parse_range(first);
            let (start2, end2) = parse_range(second);
            if contains(start1, end1, start2, end2) {
                Some(1)
            } else {
                None
            }
        })
        .sum()
}

fn contains(start1: u32, end1: u32, start2: u32, end2: u32) -> bool {
    (start1 <= start2 && end1 >= end2) || (start2 <= start1 && end2 >= end1)
}

pub fn part2(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            let (first, second) = line.split_once(',').expect("I expect a comma");
            let (start1, end1) = parse_range(first);
            let (start2, end2) = parse_range(second);
            overlap(start1, end1, start2, end2)
        })
        .filter(|v| *v)
        .count() as u32
}

fn overlap(start1: u32, end1: u32, start2: u32, end2: u32) -> bool {
    !(end1 < start2 || start1 > end2)
}

fn parse_range(s: &str) -> (u32, u32) {
    let (start, end) = s.split_once('-').expect("I expect a '-'");
    let start: u32 = start.parse().unwrap();
    let end: u32 = end.parse().unwrap();
    (start, end)
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(2, part1(&example_input()))
    }

    #[test]
    fn test_part2() {
        assert_eq!(4, part2(&example_input()))
    }
}

fn example_input() -> String {
    "\
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8"
        .to_string()
}
