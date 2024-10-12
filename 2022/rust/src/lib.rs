use std::error::Error;
use std::fs;
use std::str::FromStr;

pub fn read_one_per_line<T>(path: &str) -> Result<Vec<T>, Box<dyn Error>>
where
    T: FromStr,
{
    Ok(fs::read_to_string(path)?
        .lines()
        .filter_map(|line| line.parse().ok())
        .collect())
}

pub fn split_groups<T>(path: &str) -> Result<Vec<Vec<T>>, Box<dyn Error>>
where
    T: FromStr,
{
    Ok(fs::read_to_string(path)?
        .split("\n\n")
        .filter_map(|group| {
            let parsed: Vec<T> = group
                .lines()
                .filter_map(|line| line.trim().parse::<T>().ok())
                .collect();

            if !parsed.is_empty() {
                Some(parsed)
            } else {
                None
            }
        })
        .collect())
}
