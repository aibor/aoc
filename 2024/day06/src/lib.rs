use std::{
    collections::{HashMap, HashSet},
    fmt::{Display, Error, Formatter},
    str::FromStr,
};

use aocutils::grid::{Direction, Grid, Position};

#[derive(Debug)]
struct LabMap {
    grid: Grid,
    obstructions: HashSet<Position>,
    guard: Position,
}

impl LabMap {
    fn new(input: &str) -> Self {
        let grid = Grid::from_str(input).unwrap();

        let mut obstructions = HashSet::new();
        let mut guard = Position::default();

        for (idx, c) in input.char_indices() {
            match c {
                '#' => _ = obstructions.insert(grid.pos_from_str_idx(idx)),
                '^' => guard = grid.pos_from_str_idx(idx),
                _ => (),
            }
        }

        Self {
            grid,
            obstructions,
            guard,
        }
    }

    fn is_obstruction(&self, pos: &Position) -> bool {
        self.obstructions.contains(pos)
    }

    fn guard_path(&self) -> GuardPath<'_> {
        GuardPath {
            lab_map: self,
            dir: Direction::Up,
            pos: self.guard,
            visited: HashMap::new(),
            new_obstruction: None,
        }
    }
}

#[derive(Debug, Clone, Copy, Eq, Hash, PartialEq)]
enum Op {
    TurnRight,
    Move,
}

impl Op {
    #[must_use]
    fn is_move(&self) -> bool {
        matches!(self, Self::Move)
    }
}

#[derive(Debug, Clone)]
struct GuardPath<'a> {
    lab_map: &'a LabMap,
    pos: Position,
    dir: Direction,
    visited: HashMap<Position, HashSet<Direction>>,
    new_obstruction: Option<Position>,
}

impl<'a> GuardPath<'a> {
    fn mark_visited(&mut self) -> bool {
        let v = self.visited.entry(self.pos).or_default();
        v.insert(self.dir)
    }

    fn is_visited(&self, pos: &Position) -> bool {
        self.visited.contains_key(pos)
    }

    fn is_visited_dir(&self, pos: &Position, dir: &Direction) -> bool {
        let Some(v) = self.visited.get(pos) else {
            return false;
        };
        v.contains(dir)
    }

    fn is_obstruction(&self, pos: &Position) -> bool {
        self.lab_map.is_obstruction(pos) || self.is_test_obstruction(pos)
    }

    fn is_test_obstruction(&self, pos: &Position) -> bool {
        match self.new_obstruction {
            Some(p) => p == *pos,
            _ => false,
        }
    }

    fn pos_char(&self, pos: &Position) -> char {
        if self.is_test_obstruction(pos) {
            'O'
        } else if &self.lab_map.guard == pos {
            '^'
        } else if self.lab_map.is_obstruction(pos) {
            '#'
        } else if self.is_visited(pos) {
            let e = self.visited.get(pos).unwrap();
            if !e.contains(&Direction::Up) && !e.contains(&Direction::Down) {
                '-'
            } else if !e.contains(&Direction::Right) && !e.contains(&Direction::Left) {
                '|'
            } else {
                '+'
            }
        } else {
            '.'
        }
    }
}

impl<'a> Iterator for GuardPath<'a> {
    type Item = Op;

    fn next(&mut self) -> Option<Self::Item> {
        let next = self.pos.move_dir(&self.dir);
        if !self.lab_map.grid.valid_pos(&next) {
            None
        } else if self.is_obstruction(&next) {
            self.dir = self.dir.rotate_right();
            Some(Op::TurnRight)
        } else {
            self.pos = next;
            Some(Op::Move)
        }
    }
}

impl<'a> Display for GuardPath<'a> {
    fn fmt(&self, f: &mut Formatter<'_>) -> Result<(), Error> {
        for pos in self.lab_map.grid.line_iter() {
            if pos.x == 0 {
                writeln!(f)?;
            }
            write!(f, "{}", self.pos_char(&pos))?;
        }
        Ok(())
    }
}

pub fn part1(input: &str) -> usize {
    let lab_map = LabMap::new(input);
    let mut iter = lab_map.guard_path();

    while iter.next().is_some() {
        iter.mark_visited();
    }

    iter.visited.len()
}

pub fn part2(input: &str) -> usize {
    let lab_map = LabMap::new(input);
    let mut iter = lab_map.guard_path();
    let mut new_obstructions = HashSet::new();

    while let Some(op) = iter.next() {
        iter.mark_visited();

        if !op.is_move() {
            continue;
        };

        let next = iter.pos.move_dir(&iter.dir);
        if lab_map.guard == next || iter.is_visited(&next) || iter.is_obstruction(&next) {
            continue;
        }

        let mut i = iter.clone();
        i.new_obstruction = Some(next);

        while i.next().is_some() {
            if i.is_visited_dir(&i.pos, &i.dir) {
                new_obstructions.insert(next);
                break;
            }
            i.mark_visited();
        }
    }

    new_obstructions.len()
}

// day02 example correct, with input wrong. :(
aocutils::assert_parts!(41, 5067, 6, 1754);
