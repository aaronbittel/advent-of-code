use std::error::Error;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", part1("./data/input2.txt")?);
    println!("Part2: {}", part2("./data/input2.txt")?);

    Ok(())
}

#[derive(Debug, PartialEq, Clone)]
enum Move {
    Rock = 1,
    Paper = 2,
    Scissors = 3,
}

#[derive(Debug)]
enum Outcome {
    Win = 6,
    Loss = 0,
    Draw = 3,
}

impl Outcome {
    fn from(s: &str) -> Self {
        match s {
            "X" => Outcome::Loss,
            "Y" => Outcome::Draw,
            "Z" => Outcome::Win,
            _ => panic!("unknown outcome"),
        }
    }
}

impl Move {
    fn same(m: &Move) -> Self {
        match m {
            Move::Rock => Move::Rock,
            Move::Paper => Move::Paper,
            Move::Scissors => Move::Scissors,
        }
    }

    fn from(s: &str) -> Self {
        match s {
            "A" | "X" => Move::Rock,
            "B" | "Y" => Move::Paper,
            "C" | "Z" => Move::Scissors,
            _ => panic!("unknown play"),
        }
    }

    fn play(&self, other: &Move) -> Outcome {
        if self == other {
            return Outcome::Draw;
        }

        match (self, other) {
            (Move::Rock, Move::Scissors)
            | (Move::Scissors, Move::Paper)
            | (Move::Paper, Move::Rock) => Outcome::Win,
            _ => Outcome::Loss,
        }
    }
}

pub fn part1(path: &str) -> Result<u32, Box<dyn Error>> {
    let input = aoc::read_one_per_line::<String>(path)?;
    input
        .into_iter()
        .map(|line| {
            line.split_whitespace()
                .map(Move::from)
                .collect::<Vec<Move>>()
        })
        .map(|game| {
            let opp = &game[0];
            let me = &game[1];
            (me.play(opp) as u32) + (me.clone() as u32)
        })
        .fold(Ok(0), |acc, x| Ok(acc.unwrap() + x))
}

pub fn part2(path: &str) -> Result<u32, Box<dyn Error>> {
    let mut result: u32 = 0;

    for line in aoc::read_one_per_line::<String>(path)? {
        let parts: Vec<&str> = line.split_whitespace().collect();
        let opp = Move::from(parts[0]);
        let outcome = Outcome::from(parts[1]);
        let me = what_to_play(&opp, &outcome);
        result += (outcome as u32) + (me as u32);
    }
    Ok(result)
}

fn what_to_play(opp: &Move, outcome: &Outcome) -> Move {
    match outcome {
        Outcome::Draw => Move::same(opp),
        Outcome::Win => match opp {
            Move::Rock => Move::Paper,
            Move::Paper => Move::Scissors,
            Move::Scissors => Move::Rock,
        },
        Outcome::Loss => match opp {
            Move::Rock => Move::Scissors,
            Move::Paper => Move::Rock,
            Move::Scissors => Move::Paper,
        },
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(15, part1("./data/example2.txt").unwrap())
    }

    #[test]
    fn part2_test() {
        assert_eq!(12, part2("./data/example2.txt").unwrap())
    }
}
