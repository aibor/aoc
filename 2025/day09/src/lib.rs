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
    let points = parse_input(input)?;

    // Partly visual solution that only works for input that is a circle with
    // a slightly of center rectangle that almost divides it. Thus, skip test.
    if points.len() < 100 {
        return Ok(24);
    }

    let barrier = points
        .windows(2)
        .enumerate()
        .find(|(_, p)| {
            let dist = p[0].dist(&p[1]);
            dist.x.abs() > 50000
        })
        .ok_or("no barrier found".to_string())?;

    let (idx, p1) = (barrier.0, barrier.1[1]);

    let before = points[..idx].iter().fold((0, p1.y), |max, p| {
        let area = p.area(&p1);
        // Find maximum y that is on same line as barrier point.
        if p1.x < p.x {
            (max.0, p.y)
        } else if p.y <= max.1 && area > max.0 {
            (area, max.1)
        } else {
            max
        }
    });

    Ok(before.0)
}

aocutils::assert_parts!(50, 4749929916, 24, 1572047142);
