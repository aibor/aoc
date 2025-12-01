use std::num::ParseIntError;

fn parse_line(line: &str) -> Result<isize, String> {
    let (dir, dist) = line.split_at_checked(1).ok_or("not a rotation")?;
    let dist: isize = dist.parse().map_err(|e| ParseIntError::to_string(&e))?;
    match dir {
        "L" => Ok(-dist),
        "R" => Ok(dist),
        _ => Err("invalid direction {dir}".to_string()),
    }
}

fn parse_input(input: &str) -> Result<Vec<isize>, String> {
    input.lines().map(parse_line).collect::<Result<Vec<_>, _>>()
}

pub fn part1(input: &str) -> Result<usize, String> {
    let rotations = parse_input(input)?;
    let mut dial: isize = 50;

    let nulls = rotations.iter().fold(0, |null, r| {
        dial = (dial + *r).rem_euclid(100);
        match dial {
            0 => null + 1,
            _ => null,
        }
    });

    Ok(nulls)
}

pub fn part2(input: &str) -> Result<isize, String> {
    let rotations = parse_input(input)?;
    let mut dial: isize = 50;

    let nulls = rotations.iter().fold(0, |mut null, r| {
        let new_dial = dial + r;
        null += new_dial.abs() / 100;
        if dial > 0 && new_dial <= 0 {
            null += 1;
        }
        dial = new_dial.rem_euclid(100);
        null
    });

    Ok(nulls)
}

aocutils::assert_parts!(3, 1102, 6, 6175);
