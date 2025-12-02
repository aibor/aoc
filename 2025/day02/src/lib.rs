use std::{num::ParseIntError, ops::RangeInclusive};

type IDRange = RangeInclusive<usize>;

fn parse_id_range(s: &str) -> Result<IDRange, String> {
    let (left, right) = s.split_once("-").ok_or("no delimiter found".to_string())?;
    let start = left.parse().map_err(|e| ParseIntError::to_string(&e))?;
    let end = right.parse().map_err(|e| ParseIntError::to_string(&e))?;
    Ok(start..=end)
}

fn parse_input(input: &str) -> Result<Vec<IDRange>, String> {
    input.trim().split(",").map(parse_id_range).collect()
}

fn repeated_two(id: &usize) -> bool {
    let s = id.to_string();
    if !s.len().is_multiple_of(2) {
        return false;
    }
    let (left, right) = s.split_at(s.len() / 2);
    left == right
}

fn repeated(id: &usize) -> bool {
    let s = id.to_string();
    let d = s.repeat(2);
    d[1..(d.len() - 1)].contains(&s)
}

pub fn part1(input: &str) -> Result<usize, String> {
    let id_ranges = parse_input(input)?;

    let invalid_ids = id_ranges
        .into_iter()
        .map(|r| r.filter(repeated_two).sum::<usize>())
        .sum();

    Ok(invalid_ids)
}

pub fn part2(input: &str) -> Result<usize, String> {
    let id_ranges = parse_input(input)?;

    let invalid_ids = id_ranges
        .into_iter()
        .map(|r| r.filter(repeated).sum::<usize>())
        .sum();

    Ok(invalid_ids)
}

aocutils::assert_parts!(1227775554, 17077011375, 4174379265, 36037497037);
