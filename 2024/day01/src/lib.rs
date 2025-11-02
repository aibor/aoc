use std::collections::HashMap;

fn read_lists(input: &str) -> (Vec<u32>, Vec<u32>) {
    let mut left: Vec<u32> = Vec::new();
    let mut right: Vec<u32> = Vec::new();

    let parts = input.split_whitespace();
    for (i, part) in parts.enumerate() {
        let f: u32 = part.parse().expect("parse num");
        match i % 2 == 0 {
            true => left.push(f),
            false => right.push(f),
        }
    }

    (left, right)
}

pub fn part1(input: &str) -> u32 {
    let (mut left, mut right) = read_lists(input);
    left.sort();
    right.sort();

    left.iter()
        .zip(right.iter())
        .fold(0, |acc, (l, r)| acc + l.abs_diff(*r))
}

pub fn part2(input: &str) -> u32 {
    let (left, right) = read_lists(input);
    let mut right_count = HashMap::new();

    for r in right {
        let count = right_count.entry(r).or_insert(0);
        *count += 1;
    }

    left.iter().fold(0, |acc, l| {
        let count = right_count.get(l).copied().unwrap_or(0);
        acc + l * count
    })
}

aocutils::assert_parts!(11, 1603498, 31, 25574739);
