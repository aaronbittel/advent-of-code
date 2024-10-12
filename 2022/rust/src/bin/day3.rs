use std::collections::HashSet;
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", part1("./data/input3.txt")?);
    println!("Part2: {}", part2("./data/input3.txt")?);

    Ok(())
}

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

pub fn part1(path: &str) -> Result<u32, Box<dyn Error>> {
    Ok(aoc::read_one_per_line::<String>(path)?
        .into_iter()
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
        .sum())
}

pub fn part2(path: &str) -> Result<u32, Box<dyn Error>> {
    Ok(aoc::read_one_per_line::<String>(path)?
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
        .sum())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(157, part1("./data/example3.txt").unwrap())
    }

    #[test]
    fn test_part2() {
        assert_eq!(70, part2("./data/example3.txt").unwrap())
    }
}
