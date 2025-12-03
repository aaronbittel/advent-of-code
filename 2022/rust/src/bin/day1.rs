use std::cmp;
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", part1("./data/input1.txt")?);
    println!("Part2: {}", part2("./data/input1.txt")?);

    Ok(())
}

pub fn part1(path: &str) -> Result<i32, Box<dyn Error>> {
    let elves = aoc::split_groups::<i32>(path)?;
    let mut max_elf: i32 = 0;

    for elf in elves {
        max_elf = cmp::max(max_elf, elf.iter().sum());
    }

    Ok(max_elf)
}

pub fn part2(path: &str) -> Result<i32, Box<dyn Error>> {
    let mut supplies: Vec<i32> = aoc::split_groups::<i32>(path)?
        .iter()
        .map(|x| x.iter().sum())
        .collect();
    supplies.sort_by(|a, b| b.cmp(a));
    Ok(supplies.iter().take(3).sum())
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(24_000, part1("./data/example1.txt").unwrap());
    }

    #[test]
    fn part2_test() {
        assert_eq!(45_000, part2("./data/example1.txt").unwrap());
    }
}
