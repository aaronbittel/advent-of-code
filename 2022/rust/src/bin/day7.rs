use std::error::Error;

enum FileSystemEntity {
    File(File),
    Directory(Directory),
}

struct File {
    name: String,
    size: usize,
}

struct Directory {
    parent: Option<Box<Directory>>,
    children: Option<Vec<FileSystemEntity>>,
}

impl Directory {
    fn new() -> Self {
        Directory {
            parent: None,
            children: None,
        }
    }

    fn cd(&mut self, name: &str) {
        if name == "/" {
            let cur = self;
            while cur.parent.is_some() {
                cur = cur.parent.unwrap();
            }
        } else {
        }
    }
}

fn main() -> Result<(), Box<dyn Error>> {
    println!("Part1: {}", part1("./data/input7.txt")?);
    println!("Part2: {}", part2("./data/input7.txt")?);

    Ok(())
}

fn parse(path: &str) -> Result<Directory, Box<dyn Error>> {
    let root = Directory::new();
    let mut cur_dir = root;
    for op in &aoc::read_one_per_line::<String>(path)? {
        if op.starts_with("$ cd ") {
            cur_dir.cd(&op[5..]);
            // root.cd(op[5..]);
        }
    }
    todo!()
}

pub fn part1(path: &str) -> Result<u32, Box<dyn Error>> {
    let _ = parse(path);
    Ok(0)
}

pub fn part2(path: &str) -> Result<u32, Box<dyn Error>> {
    Ok(0)
}

#[cfg(test)]
mod tests {

    use super::*;

    #[test]
    fn test_part1() {
        assert_eq!(95437, part1("./data/example7.txt").unwrap())
    }

    #[test]
    #[ignore]
    fn test_part2() {
        assert_eq!(19, part2("./data/example7.txt").unwrap())
    }
}
