use std::str::FromStr;

#[derive(Debug, Copy, Clone, PartialEq)]
enum Operator {
    Add,
    Multiply,
}

impl FromStr for Operator {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        match s {
            "+" => Ok(Operator::Add),
            "*" => Ok(Operator::Multiply),
            _ => Err("invalid operator {s}".to_string()),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
struct Problem {
    operands: Vec<usize>,
    operator: Operator,
}

impl Problem {
    fn solve(&self) -> usize {
        let initial = match self.operator {
            Operator::Add => 0,
            Operator::Multiply => 1,
        };

        self.operands
            .iter()
            .fold(initial, |acc, operand| match self.operator {
                Operator::Add => acc + operand,
                Operator::Multiply => acc * operand,
            })
    }
}

impl FromStr for Problem {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        Ok(Self {
            operands: Vec::new(),
            operator: Operator::from_str(s)?,
        })
    }
}

fn parse_horizontal(input: &str) -> Result<Vec<Problem>, String> {
    let mut lines = input.lines().rev();
    let operators = lines.next().ok_or("missing operators".to_string())?;

    let mut problems = operators
        .split_whitespace()
        .map(Problem::from_str)
        .collect::<Result<Vec<Problem>, String>>()?;

    for line in lines {
        for (idx, s) in line.split_whitespace().enumerate() {
            let operand = s.parse::<usize>().map_err(|e| e.to_string())?;
            problems[idx].operands.push(operand);
        }
    }

    Ok(problems)
}

fn parse_vertical(input: &str) -> Result<Vec<Problem>, String> {
    let mut lines = input.lines().rev();
    let operators = lines.next().ok_or("missing operators".to_string())?;

    let mut problems = operators
        .split_whitespace()
        .map(Problem::from_str)
        .collect::<Result<Vec<Problem>, String>>()?;

    let mut idx = problems.len() - 1;
    let lines: Vec<&[u8]> = lines.map(str::as_bytes).collect();

    for i in (0..lines[0].len()).rev() {
        let s: String = (0..lines.len())
            .rev()
            .map(|j| lines[j][i] as char)
            .collect();

        let s = s.trim();
        if s.is_empty() {
            idx -= 1;
            continue;
        }

        let operand = s.parse::<usize>().map_err(|e| e.to_string())?;
        problems[idx].operands.push(operand);
    }

    Ok(problems)
}

pub fn part1(input: &str) -> Result<usize, String> {
    let problems = parse_horizontal(input)?;

    let total = problems.iter().map(Problem::solve).sum();

    Ok(total)
}

pub fn part2(input: &str) -> Result<usize, String> {
    let problems = parse_vertical(input)?;

    let total = problems.iter().map(Problem::solve).sum();

    Ok(total)
}

aocutils::assert_parts!(4277556, 5060053676136, 3263827, 9695042567249);
