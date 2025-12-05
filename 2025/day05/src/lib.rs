use std::{ops::Range, str::FromStr};

use aocutils::range::merged_ranges;

struct Ingredients {
    fresh: Vec<Range<usize>>,
    available: Vec<usize>,
}

impl FromStr for Ingredients {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut lines = s.lines();
        let parse = |s: &str| s.parse::<usize>().map_err(|e| e.to_string());

        let fresh = (&mut lines)
            .take_while(|line| !line.is_empty())
            .map(|line| {
                let (start, end) = line.split_once("-").ok_or("no delim".to_string())?;
                Ok(parse(start)?..(parse(end)? + 1))
            })
            .collect::<Result<Vec<Range<usize>>, String>>()?;

        let available = lines
            .map(|line| line.parse::<usize>().map_err(|e| e.to_string()))
            .collect::<Result<Vec<usize>, String>>()?;

        Ok(Ingredients { fresh, available })
    }
}

impl Ingredients {
    fn is_fresh(&self, id: &usize) -> bool {
        self.fresh.iter().any(|r| r.contains(id))
    }
}

pub fn part1(input: &str) -> Result<usize, String> {
    let ingredients = Ingredients::from_str(input)?;

    let fresh = ingredients
        .available
        .iter()
        .filter(|id| ingredients.is_fresh(id))
        .count();

    Ok(fresh)
}

pub fn part2(input: &str) -> Result<usize, String> {
    let ingredients = Ingredients::from_str(input)?;

    let fresh = merged_ranges(&ingredients.fresh).map(|r| r.len()).sum();

    Ok(fresh)
}

aocutils::assert_parts!(3, 613, 14, 336495597913098);
