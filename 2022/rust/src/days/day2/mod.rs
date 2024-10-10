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

pub fn part1(input: &str) -> u32 {
    input
        .lines()
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
        .fold(0, |acc, x| acc + x)
}

pub fn part2(input: &str) -> u32 {
    let lines: Vec<&str> = input.lines().collect();

    let mut result: u32 = 0;

    for line in lines {
        let parts: Vec<&str> = line.split_whitespace().collect();
        let opp = Move::from(parts[0]);
        let outcome = Outcome::from(parts[1]);
        let me = what_to_play(&opp, &outcome);
        result += (outcome as u32) + (me as u32);
    }
    result
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

fn example_input() -> String {
    "\
A Y
B X
C Z"
    .to_string()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn part1_test() {
        assert_eq!(15, part1(&example_input()));
    }

    #[test]
    fn part2_test() {
        assert_eq!(12, part2(&example_input()));
    }
}
