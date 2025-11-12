use std::{num::ParseIntError, str::FromStr};

use itertools::{Itertools, repeat_n};

#[derive(Debug)]
struct Equation {
    test_value: usize,
    numbers: Vec<usize>,
}

impl FromStr for Equation {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parse_err = |e| ParseIntError::to_string(&e);

        let (val_s, nums_s) = s.split_once(": ").ok_or("invalid input")?;
        let test_value: usize = val_s.parse().map_err(parse_err)?;
        let numbers: Result<Vec<usize>, String> = nums_s
            .split(" ")
            .map(|n| n.parse::<usize>().map_err(parse_err))
            .collect();

        Ok(Self {
            test_value,
            numbers: numbers?,
        })
    }
}

#[derive(Debug)]
enum Op {
    Add,
    Mul,
    Concat,
}

impl Op {
    fn calc(&self, mut a: usize, b: usize) -> usize {
        match self {
            Op::Add => a + b,
            Op::Mul => a * b,
            Op::Concat => {
                let mut c = b;
                while c > 0 {
                    a *= 10;
                    c /= 10;
                }
                a + b
            }
        }
    }
}

impl Equation {
    fn solveable_with(&self, op_set: &[Op]) -> bool {
        let n = self.numbers.len() - 1;
        let mut ops = repeat_n(op_set, n).multi_cartesian_product();
        ops.any(|op| self.try_solve(&op))
    }

    fn try_solve(&self, op: &[&Op]) -> bool {
        let mut o = op.iter();
        let mut acc = 0;

        for b in self.numbers.iter() {
            acc = match acc {
                v if v > self.test_value => return false,
                0 => *b,
                _ => o.next().expect("ops exhausted").calc(acc, *b),
            }
        }

        acc == self.test_value
    }
}

fn parse_input(input: &str) -> Vec<Equation> {
    input
        .lines()
        .map(|l| l.parse().expect("failed to parse line: {l}"))
        .collect()
}

pub fn part1(input: &str) -> usize {
    parse_input(input)
        .iter()
        .filter(|e| e.solveable_with(&[Op::Add, Op::Mul]))
        .fold(0, |acc, e| acc + e.test_value)
}

pub fn part2(input: &str) -> usize {
    parse_input(input)
        .iter()
        .filter(|e| e.solveable_with(&[Op::Add, Op::Mul, Op::Concat]))
        .fold(0, |acc, e| acc + e.test_value)
}

aocutils::assert_parts!(3749, 1430271835320, 11387, 456565678667482);
