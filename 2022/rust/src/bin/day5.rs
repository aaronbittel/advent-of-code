use std::collections::VecDeque;
use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", part1("./data/input5.txt")?);
    println!("Part2: {}", part2("./data/input5.txt")?);

    Ok(())
}

pub fn part1(path: &str) -> Result<String, Box<dyn Error>> {
    let mut ship = parse(path);
    for inst in &ship.instructions {
        for _ in 0..inst.count {
            let crate_supply = ship.crates[inst.from].pop_back().unwrap();
            ship.crates[inst.to].push_back(crate_supply);
        }
    }
    Ok(ship.answer())
}

pub fn part2(path: &str) -> Result<String, Box<dyn Error>> {
    let mut ship = parse(path);
    for inst in &ship.instructions {
        let split_point = ship.crates[inst.from].len() - inst.count;
        let mut crates = ship.crates[inst.from].split_off(split_point);
        ship.crates[inst.to].append(&mut crates);
    }
    Ok(ship.answer())
}

#[derive(Debug)]
struct Ship {
    crates: Vec<VecDeque<char>>,
    instructions: Vec<Instruction>,
}

impl Ship {
    fn answer(&self) -> String {
        self.crates
            .iter()
            .map(|c| *c.back().unwrap())
            .collect::<String>()
    }
}

#[derive(Debug)]
struct Instruction {
    count: usize,
    from: usize,
    to: usize,
}

fn parse(path: &str) -> Ship {
    let mut ship = Ship {
        crates: Vec::new(),
        instructions: Vec::new(),
    };

    let lines: Vec<String> = aoc::read_one_per_line::<String>(path).unwrap();
    for line in lines {
        if line.starts_with("move") {
            let parts: Vec<&str> = line.split(" ").collect();
            ship.instructions.push(Instruction {
                count: parts[1].parse().unwrap(),
                from: parts[3].parse::<usize>().unwrap() - 1,
                to: parts[5].parse::<usize>().unwrap() - 1,
            });
        } else if line.contains("[") {
            line.chars()
                .enumerate()
                .filter(|(_, c)| c.is_alphabetic())
                .for_each(|(i, c)| {
                    let idx = (i - 1) / 4;
                    while ship.crates.len() < idx + 1 {
                        ship.crates.push(VecDeque::new());
                    }
                    ship.crates[idx].push_front(c)
                })
        }
    }

    ship
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(String::from("CMZ"), part1("./data/example5.txt").unwrap())
    }

    #[test]
    fn test_part2() {
        assert_eq!(String::from("MCD"), part2("./data/example5.txt").unwrap())
    }
}
