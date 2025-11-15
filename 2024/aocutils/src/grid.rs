use std::{
    cell::LazyCell,
    ops::{Add, AddAssign, Sub, SubAssign},
    str::FromStr,
};

#[derive(Debug, Clone, Copy, Eq, Hash, PartialEq)]
pub enum Direction {
    Up,
    Right,
    Down,
    Left,
}

impl Direction {
    pub fn diff(&self) -> Position {
        let (x, y) = match self {
            Self::Up => (0, 1),
            Self::Right => (1, 0),
            Self::Down => (0, -1),
            Self::Left => (-1, 0),
        };
        Position { x, y }
    }

    pub fn rotate_right(&self) -> Self {
        match self {
            Self::Up => Self::Right,
            Self::Right => Self::Down,
            Self::Down => Self::Left,
            Self::Left => Self::Up,
        }
    }
}

#[derive(Debug, Default, Clone, Copy, Eq, Hash, PartialEq)]
pub struct Position {
    pub x: isize,
    pub y: isize,
}

impl Position {
    pub fn move_dir(&self, dir: &Direction) -> Self {
        let diff = dir.diff();

        Self {
            x: self.x + diff.x,
            y: self.y + diff.y,
        }
    }

    pub fn dist(&self, other: &Self) -> Self {
        let mut x = self.x.abs_diff(other.x) as isize;
        if self.x > other.x {
            x = -x;
        }

        let mut y = self.y.abs_diff(other.y) as isize;
        if self.y > other.y {
            y = -y;
        }

        Self { x, y }
    }
}

impl Add for Position {
    type Output = Self;

    fn add(self, other: Self) -> Self::Output {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl AddAssign for Position {
    fn add_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}

impl Sub for Position {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        Self {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

impl SubAssign for Position {
    fn sub_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x - other.x,
            y: self.y - other.y,
        };
    }
}

#[derive(Debug)]
pub struct Grid {
    pub width: usize,
    pub height: usize,
}

impl FromStr for Grid {
    type Err = String;

    fn from_str(input: &str) -> Result<Self, Self::Err> {
        let width = input.find('\n').expect("no newline found");
        let height = input.lines().count();

        Ok(Self { width, height })
    }
}

impl Grid {
    pub fn pos_from_str_idx(&self, idx: usize) -> Position {
        let line = idx / (self.width + 1);
        Position {
            x: (idx % (self.width + 1)) as isize,
            y: (self.height - line - 1) as isize,
        }
    }

    pub fn valid_pos(&self, pos: &Position) -> bool {
        let rx = LazyCell::new(|| 0..self.width as isize);
        let ry = LazyCell::new(|| 0..self.height as isize);
        rx.contains(&pos.x) && ry.contains(&pos.y)
    }

    pub fn iter(&self) -> GridIterator<'_> {
        let start = Position {
            x: 0,
            y: (self.height - 1) as isize,
        };
        GridIterator {
            grid: self,
            pos: start,
            step: Position { x: 1, y: -1 },
        }
    }

    pub fn line_iter(&self) -> GridIterator<'_> {
        let start = Position {
            x: 0,
            y: (self.height - 1) as isize,
        };
        GridIterator {
            grid: self,
            pos: start,
            step: Position { x: 1, y: -1 },
        }
    }
}

pub struct GridIterator<'a> {
    grid: &'a Grid,
    pos: Position,
    step: Position,
}

impl<'a> Iterator for GridIterator<'a> {
    type Item = Position;
    fn next(&mut self) -> Option<Self::Item> {
        let cur = self.pos;
        if cur.x == (self.grid.width - 1) as isize {
            self.pos.x = 0;
            self.pos.y += self.step.y;
        } else {
            self.pos.x += self.step.x;
        }
        if self.grid.valid_pos(&cur) {
            Some(cur)
        } else {
            None
        }
    }
}
