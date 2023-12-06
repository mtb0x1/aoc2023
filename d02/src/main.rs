use anyhow::Result;


///source of challenge is : https://adventofcode.com/2023/day/2

fn data() -> &'static str {
    include_str!("input.txt")
}

fn solve_part1_and_part2(input: &str) -> (i64, i64) {
    let mut part1 = 0;
    let mut part2 = 0;

    for l in input.lines() {
        //eprintln!("processing line n° {}",i);
        
        let (game, pulls) = match l.trim().split_once(':') {
            Some((a, b)) => (a, b),
            _ => panic!(),
        };

        let game_id: i64 = match game.strip_prefix("Game") {
            Some(val) => val.trim().parse().unwrap(),
            _ => panic!(),
        };

        let colors = ["red", "green", "blue"];
        let mut color_bounds = [0, 0, 0];
        for pull in pulls.split(';') {
            for cube in pull.split(',') {
                
                let (count, color) = match cube.trim().split_once(' ') {
                    Some((a, b)) => (a, b),
                    _ => panic!(),
                };

                let color_idx = match colors.iter().position(|c| *c == color) {
                    Some(val) => val,
                    _ => panic!(),
                };
                
                color_bounds[color_idx] = color_bounds[color_idx].max(count.parse().unwrap());
            }
        }

        let possible = color_bounds[0] <= 12 && color_bounds[1] <= 13 && color_bounds[2] <= 14;
        part1 += if possible { game_id } else { 0 };
        part2 += color_bounds.iter().product::<i64>();
    }
    (part1, part2)
}

pub fn main() -> Result<()> {
    let data = data();
    let (p1, p2) = solve_part1_and_part2(data);
    println!("part 1 = {p1}\npart 2 = {p2}");
    Ok(())
}
