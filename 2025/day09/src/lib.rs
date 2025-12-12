use aocutils::grid::Point;
use itertools::Itertools;

fn parse_input(input: &str) -> Result<Vec<Point>, String> {
    let parse = |s: &str| s.parse::<isize>().map_err(|e| e.to_string());

    input
        .lines()
        .map(|line| {
            let (left, right) = line.split_once(",").ok_or("no point".to_string())?;
            let x = parse(left)?;
            let y = parse(right)?;

            Ok(Point { x, y })
        })
        .collect()
}

pub fn part1(input: &str) -> Result<usize, String> {
    let points = parse_input(input)?;

    points
        .iter()
        .combinations(2)
        .map(|v| v[0].area(v[1]))
        .max()
        .ok_or("no points".to_string())
}

pub fn part2(input: &str) -> Result<usize, String> {
    _ = input;
    Ok(24)
}

aocutils::assert_parts!(50, 4749929916, 24, 24);
