use std::{
    cell::LazyCell,
    fmt::Display,
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
    pub fn diff(&self) -> Point {
        let (x, y) = match self {
            Self::Up => (0, 1),
            Self::Right => (1, 0),
            Self::Down => (0, -1),
            Self::Left => (-1, 0),
        };
        Point { x, y }
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
pub struct Point {
    pub x: isize,
    pub y: isize,
}

impl Point {
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

    pub fn area(&self, other: &Self) -> usize {
        let x = self.x.abs_diff(other.x) + 1;
        let y = self.y.abs_diff(other.y) + 1;
        x * y
    }
}

impl Add for Point {
    type Output = Self;

    fn add(self, other: Self) -> Self::Output {
        Self {
            x: self.x + other.x,
            y: self.y + other.y,
        }
    }
}

impl AddAssign for Point {
    fn add_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x + other.x,
            y: self.y + other.y,
        };
    }
}

impl Sub for Point {
    type Output = Self;

    fn sub(self, other: Self) -> Self::Output {
        Self {
            x: self.x - other.x,
            y: self.y - other.y,
        }
    }
}

impl SubAssign for Point {
    fn sub_assign(&mut self, other: Self) {
        *self = Self {
            x: self.x - other.x,
            y: self.y - other.y,
        };
    }
}

impl Display for Point {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
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
    pub fn point_from_str_idx(&self, idx: usize) -> Point {
        let line = idx / (self.width + 1);
        Point {
            x: (idx % (self.width + 1)) as isize,
            y: (self.height - line - 1) as isize,
        }
    }

    pub fn valid(&self, pos: &Point) -> bool {
        let rx = LazyCell::new(|| 0..self.width as isize);
        let ry = LazyCell::new(|| 0..self.height as isize);
        rx.contains(&pos.x) && ry.contains(&pos.y)
    }

    pub fn iter(&self) -> GridIterator<'_> {
        GridIterator {
            grid: self,
            pos: Point { x: 0, y: 0 },
            step: Point { x: 1, y: 1 },
        }
    }

    pub fn line_iter(&self) -> GridIterator<'_> {
        let start = Point {
            x: 0,
            y: (self.height - 1) as isize,
        };
        GridIterator {
            grid: self,
            pos: start,
            step: Point { x: 1, y: -1 },
        }
    }
}

pub struct GridIterator<'a> {
    grid: &'a Grid,
    pos: Point,
    step: Point,
}

impl<'a> Iterator for GridIterator<'a> {
    type Item = Point;

    fn next(&mut self) -> Option<Self::Item> {
        let cur = self.pos;
        if cur.x == (self.grid.width - 1) as isize {
            self.pos.x = 0;
            self.pos.y += self.step.y;
        } else {
            self.pos.x += self.step.x;
        }
        self.grid.valid(&cur).then_some(cur)
    }
}
