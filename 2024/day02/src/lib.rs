fn report(input: &str) -> Vec<i32> {
    input
        .split_whitespace()
        .map(|s| s.parse::<i32>().unwrap())
        .collect()
}

fn increments(a: &i32, b: &i32) -> Option<bool> {
    match b - a {
        1..=3 => Some(true),
        -3..=-1 => Some(false),
        _ => None,
    }
}

fn is_safe(report: &[i32]) -> Option<()> {
    let mut levels = report.windows(2);
    let init_levels = levels.next()?;
    let report_inc = increments(&init_levels[0], &init_levels[1])?;

    for level_pair in levels {
        if report_inc != increments(&level_pair[0], &level_pair[1])? {
            return None;
        }
    }

    Some(())
}

fn is_safe_dampened(report: Vec<i32>) -> Option<()> {
    for (i, _) in report.iter().enumerate() {
        let mut k = report.clone();
        k.remove(i);
        if let Some(ret) = is_safe(&k) {
            return Some(ret);
        }
    }
    None
}

pub fn part1(input: &str) -> usize {
    input
        .lines()
        .filter(|l| is_safe(&report(l)).is_some())
        .count()
}

pub fn part2(input: &str) -> usize {
    input
        .lines()
        .filter(|l| {
            let r = report(l);
            is_safe(&r).is_some() || is_safe_dampened(r).is_some()
        })
        .count()
}

aocutils::assert_parts!(2, 526, 4, 566);
