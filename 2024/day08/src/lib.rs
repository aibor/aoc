use std::{
    collections::{HashMap, HashSet},
    str::FromStr,
};

use aocutils::grid::{Grid, Position};

struct Map {
    grid: Grid,
    antennas: HashMap<char, Vec<Position>>,
}

impl FromStr for Map {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let grid: Grid = s.parse()?;
        let mut antennas = HashMap::new();

        let antennas_iter = s
            .char_indices()
            .filter(|(_, c)| c.is_ascii_alphanumeric())
            .map(|(idx, c)| (c, grid.pos_from_str_idx(idx)));

        for (freq, pos) in antennas_iter {
            let positions: &mut Vec<Position> = antennas.entry(freq).or_default();
            positions.push(pos);
        }

        Ok(Self { grid, antennas })
    }
}

impl Map {
    fn antenna_pairs(&self) -> impl Iterator<Item = (Position, Position)> {
        self.antennas.values().flat_map(|positions| {
            positions
                .iter()
                .enumerate()
                .flat_map(|(idx, pos1)| positions.iter().skip(idx + 1).map(|pos2| (*pos1, *pos2)))
        })
    }
}

pub fn part1(input: &str) -> usize {
    let map: Map = input.parse().expect("failed to parse map");
    let mut antinodes = HashSet::new();

    let mut insert = |pos| {
        if map.grid.valid_pos(&pos) {
            antinodes.insert(pos);
        }
    };

    for (pos1, pos2) in map.antenna_pairs() {
        let dist = pos1.dist(&pos2);
        insert(pos1 - dist);
        insert(pos2 + dist);
    }

    antinodes.len()
}

pub fn part2(input: &str) -> usize {
    let map: Map = input.parse().expect("failed to parse map");
    let mut antinodes = HashSet::new();

    let mut insert = |pos| -> bool {
        if !map.grid.valid_pos(&pos) {
            return false;
        }
        antinodes.insert(pos);
        true
    };

    for (mut pos1, mut pos2) in map.antenna_pairs() {
        let dist = pos1.dist(&pos2);

        while insert(pos1) {
            pos1 -= dist;
        }

        while insert(pos2) {
            pos2 += dist
        }
    }

    antinodes.len()
}

aocutils::assert_parts!(14, 394, 34, 1277);
