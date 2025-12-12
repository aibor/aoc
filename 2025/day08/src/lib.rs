use itertools::Itertools;

#[derive(PartialEq, Eq, Default, Copy, Clone, Debug, Hash, PartialOrd, Ord)]
struct Point(usize, usize, usize);

impl Point {
    fn dist(&self, other: &Point) -> usize {
        ((other.0.abs_diff(self.0)).pow(2)
            + (other.1.abs_diff(self.1)).pow(2)
            + (other.2.abs_diff(self.2)).pow(2))
        .isqrt()
    }
}

fn parse_input(input: &str) -> Result<Vec<Point>, String> {
    input
        .lines()
        .map(|line| {
            let parts: Vec<usize> = line
                .splitn(3, ",")
                .map(|c| c.parse::<usize>().map_err(|e| e.to_string()))
                .collect::<Result<Vec<usize>, String>>()?;
            Ok(Point(parts[0], parts[1], parts[2]))
        })
        .collect()
}

fn distances(points: &[Point]) -> Vec<(&Point, &Point, usize)> {
    let mut dists: Vec<(&Point, &Point, usize)> = points
        .iter()
        .combinations(2)
        .map(|v| {
            let (p, q) = (v[0], v[1]);
            let dist = p.dist(q);
            (p, q, dist)
        })
        .collect();
    dists.sort_by_key(|a| a.2);
    dists
}

fn merge(circuits: &mut Vec<Vec<&Point>>, a: usize, b: usize) -> bool {
    if a == b {
        return false;
    }
    let (a, b) = if a > b { (b, a) } else { (a, b) };
    let c = circuits.remove(b);
    circuits[a].extend_from_slice(&c);
    true
}

pub fn part1(input: &str) -> Result<usize, String> {
    let points = parse_input(input)?;
    let max_conns = if points.len() == 20 { 10 } else { 1000 };

    let mut circuits: Vec<Vec<&Point>> = Vec::with_capacity(max_conns);
    for n in distances(&points)[..max_conns].iter() {
        match (
            circuits.iter().enumerate().find(|(_, c)| c.contains(&n.0)),
            circuits.iter().enumerate().find(|(_, c)| c.contains(&n.1)),
        ) {
            (Some((a, _)), Some((b, _))) => _ = merge(&mut circuits, a, b),
            (Some((a, _)), None) => circuits[a].push(n.1),
            (None, Some((b, _))) => circuits[b].push(n.0),
            (None, None) => circuits.push(vec![n.0, n.1]),
        }
    }

    circuits.sort_by_key(|a| a.len());

    Ok(circuits.iter().rev().take(3).map(Vec::len).product())
}

pub fn part2(input: &str) -> Result<usize, String> {
    let points = parse_input(input)?;

    let mut circuits: Vec<Vec<&Point>> = Vec::with_capacity(points.len());
    circuits.extend(points.iter().map(|p| vec![p]));

    distances(&points)
        .iter()
        .find(|n| {
            let (Some((a, _)), Some((b, _))) = (
                circuits.iter().enumerate().find(|(_, c)| c.contains(&n.0)),
                circuits.iter().enumerate().find(|(_, c)| c.contains(&n.1)),
            ) else {
                return false;
            };
            merge(&mut circuits, a, b) && circuits.len() == 1
        })
        .map(|d| d.0.0 * d.1.0)
        .ok_or("no connection found".to_string())
}

aocutils::assert_parts!(40, 79560, 25272, 31182420);
