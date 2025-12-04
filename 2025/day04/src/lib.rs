use std::{collections::HashSet, str::FromStr};

use aocutils::grid::{Direction, Grid, Position};

fn paper_rolls(s: &str) -> Result<HashSet<Position>, String> {
    let grid = Grid::from_str(s)?;
    let paper_rolls = HashSet::from_iter(s.char_indices().filter_map(|(i, c)| match c {
        '@' => Some(grid.pos_from_str_idx(i)),
        _ => None,
    }));

    Ok(paper_rolls)
}

fn is_accessible(paper_rolls: &HashSet<Position>, paper: &Position) -> bool {
    let movements = [
        Direction::Right,
        Direction::Down,
        Direction::Left,
        Direction::Left,
        Direction::Up,
        Direction::Up,
        Direction::Right,
        Direction::Right,
    ];

    let mut adjacent_papers = 0;
    let mut pos = *paper;

    for direction in movements {
        pos = pos.move_dir(&direction);
        if paper_rolls.contains(&pos) {
            adjacent_papers += 1;
        }
        if adjacent_papers >= 4 {
            return false;
        }
    }

    true
}

pub fn part1(input: &str) -> Result<usize, String> {
    let paper_rolls = paper_rolls(input)?;

    let accessible = paper_rolls
        .iter()
        .filter(|paper| is_accessible(&paper_rolls, paper))
        .count();

    Ok(accessible)
}

pub fn part2(input: &str) -> Result<usize, String> {
    let mut paper_rolls = paper_rolls(input)?;
    let mut removable = 0;

    while let Some(accessible) = paper_rolls.clone().iter().fold(None, |acc, paper| {
        match is_accessible(&paper_rolls, paper) {
            true => {
                paper_rolls.remove(paper);
                Some(acc.unwrap_or(0) + 1)
            }
            false => acc,
        }
    }) {
        removable += accessible;
    }

    Ok(removable)
}

aocutils::assert_parts!(13, 1356, 43, 8713);
