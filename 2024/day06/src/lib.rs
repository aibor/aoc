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
    visited: HashMap<Position, [bool; 4]>,
    new_obstruction: Option<Position>,
}

impl<'a> GuardPath<'a> {
    fn mark_visited(&mut self) -> bool {
        if let Some(v) = self.visited.get_mut(&self.pos) {
            v[self.dir as usize] = true;
            return false;
        }
        let mut v = [false, false, false, false];
        v[self.dir as usize] = true;
        self.visited.insert(self.pos, v);
        true
    }

    fn is_visited_dir(&self, pos: &Position, dir: &Direction) -> bool {
        let Some(v) = self.visited.get(pos) else {
            return false;
        };
        v[*dir as usize]
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
        } else if let Some(e) = self.visited.get(pos) {
            if !e[Direction::Up as usize] && !e[Direction::Down as usize] {
                '-'
            } else if !e[Direction::Right as usize] && !e[Direction::Left as usize] {
                '|'
            } else {
                '+'
            }
        } else {
            '.'
        }
    }

    fn loops(&mut self) -> bool {
        while self.next().is_some() {
            if self.is_visited_dir(&self.pos, &self.dir) {
                return true;
            }
            self.mark_visited();
        }
        false
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
    let mut i = iter.clone();

    while let Some(op) = iter.next() {
        if iter.mark_visited() && op.is_move() {
            i.new_obstruction = Some(iter.pos);
            if i.loops() {
                new_obstructions.insert(iter.pos);
            }
        }

        i = iter.clone();
    }

    new_obstructions.len()
}

aocutils::assert_parts!(41, 5067, 6, 1793);
