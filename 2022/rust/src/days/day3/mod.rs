use std::collections::HashSet;

trait Priority {
    fn priority(&self) -> u8;
}

impl Priority for char {
    fn priority(&self) -> u8 {
        match &self {
            'a'..='z' => (*self as u8) - b'a' + 1,
            'A'..='Z' => (*self as u8) - b'A' + 27,
            _ => panic!("only letters are acceptec"),
        }
    }
}

pub fn part1(input: &str) -> u32 {
    input
        .lines()
        .map(|line| {
            let mid = line.len() / 2;
            let first = &line[..mid];
            let second = &line[mid..];
            let set: HashSet<char> = first.chars().collect();
            second
                .chars()
                .find(|c| set.contains(c))
                .expect("every line should have a common character")
                .priority() as u32
        })
        .sum()
}

pub fn part2(input: &str) -> u32 {
    input
        .lines()
        .collect::<Vec<&str>>()
        .chunks(3)
        .map(|group| {
            let common = group
                .iter()
                .map(|line| line.chars().collect::<HashSet<_>>())
                .reduce(|a, b| &a & &b)
                .expect("There should be one common in the groups");

            common
                .iter()
                .next()
                .expect("There should be 1 common")
                .priority() as u32
        })
        .sum()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(157, part1(&example_input()))
    }

    #[test]
    fn test_part2() {
        assert_eq!(70, part2(&example_input()))
    }
}

fn example_input() -> String {
    "\
vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"
        .to_string()
}
