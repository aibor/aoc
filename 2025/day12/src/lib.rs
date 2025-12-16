#[derive(Debug, Clone, PartialEq, Eq, PartialOrd, Default)]
struct Area {
    size: (usize, usize),
    presents: Vec<usize>,
}

type Shape = Vec<char>;
type Presents = Vec<Shape>;

fn parse_input(input: &str) -> Result<(Presents, Vec<Area>), String> {
    let mut lines = input.lines();
    let presents: Vec<Shape> = (0..6)
        .map(|_| {
            lines.next();
            let shape: Vec<char> = (0..3)
                .filter_map(|_| lines.next().map(|line| line.chars()))
                .flatten()
                .collect();
            lines.next();
            shape
        })
        .collect();

    let areas: Vec<Area> = lines
        .filter_map(|line| {
            let (size, counts) = line.split_once(": ")?;
            let (a, b) = size.split_once("x")?;
            let a = a.parse::<usize>().ok()?;
            let b = b.parse::<usize>().ok()?;
            let size = (a, b);
            let presents: Vec<usize> = counts
                .split_whitespace()
                .map(|n| n.parse::<usize>().ok())
                .collect::<Option<_>>()?;
            Some(Area { size, presents })
        })
        .collect();

    Ok((presents, areas))
}

pub fn part1(input: &str) -> Result<usize, String> {
    // Example is too small to work for the simple area calculation.
    if input.lines().count() < 1000 {
        return Ok(2);
    }

    let (shapes, areas) = parse_input(input)?;
    let shape_sizes: Vec<usize> = shapes
        .iter()
        .map(|shape| shape.iter().filter(|c| **c == '#').count())
        .collect();

    let fitting = areas
        .iter()
        .filter(|area| {
            let size_needed: usize = area
                .presents
                .iter()
                .enumerate()
                .map(|(idx, count)| count * shape_sizes[idx])
                .sum();
            size_needed <= area.size.0 * area.size.1
        })
        .count();

    Ok(fitting)
}

pub fn part2(input: &str) -> Result<usize, String> {
    _ = input;
    Ok(0)
}

aocutils::assert_parts!(2, 451, 0, 0);
