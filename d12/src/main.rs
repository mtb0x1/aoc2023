use anyhow::Result;
fn parse() -> &'static str {
    include_str!("input.txt")
}

fn post_parse(line: &str, springs: &mut Vec<Spring>, lens: &mut Vec<usize>) {
    let (lhs, rhs) = line.split_once(' ').unwrap();
    for spring in lhs.bytes().map(Spring::from) {
        springs.push(spring);
    }
    for len in rhs.split(',').map(|len| len.parse::<usize>().unwrap()) {
        lens.push(len);
    }
}

type Matrix<T> = Vec<Vec<T>>;

#[derive(Debug, Clone, Copy, Default, PartialEq)]
enum Spring {
    #[default]
    Ok,
    Broken,
    Unknown,
}

impl From<u8> for Spring {
    fn from(value: u8) -> Self {
        match value {
            b'.' => Self::Ok,
            b'#' => Self::Broken,
            b'?' => Self::Unknown,
            _ => panic!("invalid token"),
        }
    }
}

#[derive(Clone, Copy, PartialEq, Default)]
enum Outcome {
    #[default]
    None,
    Invalid,
    Valid(usize),
}

impl Outcome {
    fn unwrap_or(&self, value: usize) -> usize {
        if let Self::Valid(outcome) = self {
            *outcome
        } else {
            value
        }
    }
}

fn arrangements_memoized(springs: &[Spring], lens: &[usize]) -> usize {
    let mut memo: Matrix<Outcome> = vec![vec![Outcome::None; lens.len() + 1]; springs.len() + 1];
    arrangements(springs, lens, &mut memo).unwrap_or(0)
}

fn place(len: usize, springs: &[Spring], lens: &[usize], memo: &mut Matrix<Outcome>) -> Outcome {
    if len > springs.len() || springs[..len].iter().any(|spring| *spring == Spring::Ok) {
        Outcome::Invalid
    } else if len >= springs.len() {
        arrangements(&springs[len..], lens, memo)
    } else if springs[len] == Spring::Broken {
        Outcome::Invalid
    } else {
        arrangements(&springs[len + 1..], lens, memo)
    }
}

fn arrangements(springs: &[Spring], lens: &[usize], memo: &mut Matrix<Outcome>) -> Outcome {
    if let memo @ (Outcome::Valid(_) | Outcome::Invalid) = memo[springs.len()][lens.len()] {
        return memo;
    }
    let outcome = match (springs.iter().next(), lens.iter().next()) {
        (None, None) => Outcome::Valid(1),
        (None, Some(_)) => Outcome::Invalid,
        (Some(Spring::Ok), _) => arrangements(&springs[1..], lens, memo),
        (Some(Spring::Broken), None) => Outcome::Invalid,
        (Some(Spring::Broken), Some(len)) => place(*len, springs, &lens[1..], memo),
        (Some(Spring::Unknown), None) => arrangements(&springs[1..], lens, memo),
        (Some(Spring::Unknown), Some(len)) => {
            let here = place(*len, springs, &lens[1..], memo).unwrap_or(0);
            let there = arrangements(&springs[1..], lens, memo).unwrap_or(0);
            Outcome::Valid(here + there)
        }
    };
    memo[springs.len()][lens.len()] = outcome;
    outcome
}

fn expand(by: usize, springs: &mut Vec<Spring>, lens: &mut Vec<usize>) {
    let springs_len = springs.len();
    let lens_len = lens.len();
    for _ in 1..by {
        springs.push(Spring::Unknown);
        for j in 0..springs_len {
            springs.push(springs[j]);
        }
        for j in 0..lens_len {
            lens.push(lens[j]);
        }
    }
}

fn sum_arrangements(input: &str, copies: usize) -> usize {
    let mut sum = 0;
    let mut springs = Vec::new();
    let mut lens = Vec::new();
    for line in input.lines() {
        springs.clear();
        lens.clear();
        post_parse(line, &mut springs, &mut lens);
        expand(copies, &mut springs, &mut lens);
        sum += arrangements_memoized(&springs[..], &lens[..]);
    }
    sum
}

fn solve_part1(input: &str) -> usize {
    sum_arrangements(input, 1)
}

fn solve_part2(input: &str) -> usize {
    sum_arrangements(input, 5)
}

pub fn main() -> Result<()> {
    let data = parse();
    let p1 = solve_part1(data);
    let p2 = solve_part2(data);
    println!("part 1 = {p1}\npart 2 = {p2}");
    Ok(())
}