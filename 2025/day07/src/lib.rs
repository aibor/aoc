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

    fn beam<F>(&self, start: &Point, mut f: F)
    where
        F: FnMut(&Point, &(Point, Point)) -> bool,
    {
        let mut beams = VecDeque::from([*start]);

        while let Some(beam) = beams.pop_front() {
            let Some(splitter) = self.find_splitter(&beam) else {
                continue;
            };

            let next = (
                splitter.move_dir(&Direction::Left),
                splitter.move_dir(&Direction::Right),
            );

            if f(&beam, &next) {
                [next.0, next.1].iter().for_each(|pos| {
                    if !beams.contains(pos) {
                        beams.push_back(*pos);
                    }
                });
            }
        }
    }
}

pub fn part1(input: &str) -> Result<usize, String> {
    let tachyon_manifold: TachyonManifold = input.parse()?;

    let mut splitters_seen = HashSet::new();

    tachyon_manifold.beam(&tachyon_manifold.start, |_, splitter| {
        splitters_seen.insert(*splitter)
    });

    Ok(splitters_seen.len())
}

pub fn part2(input: &str) -> Result<usize, String> {
    let tachyon_manifold: TachyonManifold = input.parse()?;

    let mut beam_paths: HashMap<Point, usize> = HashMap::new();

    tachyon_manifold.beam(&tachyon_manifold.start, |&hit_from, next| {
        let mut paths = 1;
        beam_paths.entry(hit_from).and_modify(|e| {
            paths = *e;
            *e = 0;
        });

        beam_paths
            .entry(next.0)
            .and_modify(|e| *e += paths)
            .or_insert(paths);

        beam_paths
            .entry(next.1)
            .and_modify(|e| *e += paths)
            .or_insert(paths);

        true
    });

    Ok(beam_paths.values().sum())
}

aocutils::assert_parts!(21, 1581, 40, 73007003089792);
