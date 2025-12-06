use std::ops::RangeInclusive;

type IDRange = RangeInclusive<usize>;

fn parse_id_range(s: &str) -> Result<IDRange, String> {
    let (start, end) = s.split_once("-").ok_or("no delimiter found".to_string())?;
    let to_int = |num: &str| num.parse::<usize>().map_err(|e| e.to_string());

    Ok(to_int(start)?..=to_int(end)?)
}

fn parse_input(input: &str) -> Result<Vec<IDRange>, String> {
    input.trim().split(",").map(parse_id_range).collect()
}

fn repeated_two(id: &usize) -> bool {
    let num_digits = id.ilog10() + 1;
    let div = 10usize.pow(num_digits / 2);
    id / div == id % div
}

fn repeated_any(id: &usize) -> bool {
    let num_digits = id.ilog10() + 1;

    (1..=(num_digits / 2)).rev().any(|i| {
        if !num_digits.is_multiple_of(i) {
            return false;
        }

        let div = 10usize.pow(i);
        let pattern = id % div;
        let mut current = id / div;

        while current > 0 {
            if current % div != pattern {
                return false;
            }
            current /= div;
        }

        true
    })
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
        .map(|r| r.filter(repeated_any).sum::<usize>())
        .sum();

    Ok(invalid_ids)
}

aocutils::assert_parts!(1227775554, 17077011375, 4174379265, 36037497037);
