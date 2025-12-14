use std::{
    collections::{HashMap, HashSet, VecDeque},
    str::FromStr,
};

use aocutils::grid::{Direction, Grid, Point};

struct TachyonManifold {
    start: Point,
    splitters: HashSet<Point>,
}

impl FromStr for TachyonManifold {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let grid: Grid = s.parse()?;

        let mut start = Point::default();
        let mut splitters = HashSet::new();

        for (idx, c) in s.char_indices() {
            match c {
                'S' => start = grid.point_from_str_idx(idx),
                '^' => _ = splitters.insert(grid.point_from_str_idx(idx)),
                _ => (),
            }
        }

        Ok(Self { start, splitters })
    }
}

impl TachyonManifold {
    fn find_splitter(&self, start: &Point) -> Option<Point> {
        let mut beam = *start;

        while beam.y >= 0 {
            if self.splitters.contains(&beam) {
                return Some(beam);
            }
            beam = beam.move_dir(&Direction::Down);
        }

        None
    }

    fn count_paths(&self, start: &Point, seen: &mut HashMap<Point, usize>) -> usize {
        if let Some(paths) = seen.get(start) {
            return *paths;
        }

        let Some(splitter) = self.find_splitter(start) else {
            return 1;
        };

        let left = splitter.move_dir(&Direction::Left);
        let right = splitter.move_dir(&Direction::Right);

        let paths = self.count_paths(&left, seen) + self.count_paths(&right, seen);

        seen.insert(*start, paths);

        paths
    }
}

pub fn part1(input: &str) -> Result<usize, String> {
    let tachyon_manifold: TachyonManifold = input.parse()?;

    let mut beams = VecDeque::from([tachyon_manifold.start]);
    let mut splitters_seen = HashSet::new();

    while let Some(beam) = beams.pop_front() {
        if let Some(splitter) = tachyon_manifold.find_splitter(&beam)
            && splitters_seen.insert(splitter)
        {
            let left = splitter.move_dir(&Direction::Left);
            let right = splitter.move_dir(&Direction::Right);

            [left, right].iter().for_each(|pos| {
                if !beams.contains(pos) {
                    beams.push_back(*pos);
                }
            });
        }
    }

    Ok(splitters_seen.len())
}

pub fn part2(input: &str) -> Result<usize, String> {
    let tachyon_manifold: TachyonManifold = input.parse()?;

    let mut beams_seen: HashMap<Point, usize> = HashMap::new();
    let timelines = tachyon_manifold.count_paths(&tachyon_manifold.start, &mut beams_seen);

    Ok(timelines)
}

aocutils::assert_parts!(21, 1581, 40, 73007003089792);
