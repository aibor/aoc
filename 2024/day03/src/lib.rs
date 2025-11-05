use regex::Regex;

const RE: &str = r"(?x)
    mul\((?P<x>\d{1,3}),(?P<y>\d{1,3})\)
    | (?P<do>do\(\))
    | (?P<dont>don't\(\))
    ";

fn extract(caps: &regex::Captures<'_>, name: &str) -> Option<u32> {
    caps.name(name)?.as_str().parse().ok()
}

fn mul(caps: &regex::Captures<'_>) -> Option<u32> {
    let x = extract(caps, "x")?;
    let y = extract(caps, "y")?;
    Some(x * y)
}

pub fn part1(input: &str) -> u32 {
    let re = Regex::new(RE).unwrap();
    re.captures_iter(input)
        .fold(0, |sum, caps| sum + mul(&caps).unwrap_or(0))
}

pub fn part2(input: &str) -> u32 {
    let mut enabled = true;
    let re = Regex::new(RE).unwrap();
    re.captures_iter(input).fold(0, |mut sum, caps| {
        if enabled {
            if let Some(p) = mul(&caps) {
                sum += p;
            } else if caps.name("dont").is_some() {
                enabled = false;
            }
        } else if caps.name("do").is_some() {
            enabled = true;
        }
        sum
    })
}

aocutils::assert_parts!(161, 178794710, 48, 76729637);
