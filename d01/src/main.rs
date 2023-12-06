use anyhow::Result;

#[allow(unused_variables)]
fn parse() -> &'static str {
    include_str!("input.txt")
}

fn solve_part1(input: &str) -> usize {
    let result: usize = input
        .lines()
        .map(|line| {
            let first_digit = line.chars().find(|c| c.is_ascii_digit()).unwrap_or('0');
            let last_digit = line
                .chars()
                .rev()
                .find(|c| c.is_ascii_digit())
                .unwrap_or('0');
            let calibration_value = format!("{}{}", first_digit, last_digit);
            calibration_value.parse::<usize>().unwrap_or(0)
        })
        .sum();
    result
}

fn solve_part2(input: &str) -> usize {
    solve_part1(
        &input
            .replace("one", "x1x")
            .replace("two", "x2x")
            .replace("three", "x3x")
            .replace("four", "x4x")
            .replace("five", "x5x")
            .replace("six", "x6x")
            .replace("seven", "x7x")
            .replace("eight", "x8x")
            .replace("nine", "x9x"),
    )
}

pub fn main() -> Result<()> {
    let data = parse();
    let p1 = solve_part1(data);
    let p2 = solve_part2(data);
    println!("part 1 = {p1}\npart 2 = {p2}");
    Ok(())
}
