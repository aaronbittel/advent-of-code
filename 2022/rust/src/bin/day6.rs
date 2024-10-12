use std::collections::HashSet;
use std::error::Error;
use std::fs;

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", window("./data/input6.txt", 4)?);
    println!("Part2: {}", window("./data/input6.txt", 14)?);

    Ok(())
}

pub fn window(path: &str, size: usize) -> Result<usize, Box<dyn Error>> {
    let content = fs::read_to_string(path)?;
    for (i, window) in content
        .chars()
        .collect::<Vec<char>>()
        .windows(size)
        .enumerate()
    {
        let set = window.into_iter().collect::<HashSet<&char>>();
        if set.len() == size {
            return Ok(i + size);
        }
    }

    Err("Could not find window".into())
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(10, window("./data/example6.txt", 4).unwrap())
    }

    #[test]
    fn test_part2() {
        assert_eq!(29, window("./data/example6.txt", 14).unwrap())
    }
}
