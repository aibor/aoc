fn banks(input: &str) -> Vec<&[u8]> {
    input.lines().map(str::as_bytes).collect()
}

fn max_battery(batteries: &[u8]) -> (usize, u8) {
    batteries.iter().enumerate().fold(
        (0, 0),
        |max, (i, &bat)| if bat > max.1 { (i, bat) } else { max },
    )
}

fn joltage(bank: &[u8], battery_count: usize) -> Result<usize, String> {
    let mut start_idx = 0;

    let batteries: Vec<u8> = (0..battery_count)
        .rev()
        .map(|i| {
            let end_idx = bank.len() - i;
            let battery = max_battery(&bank[start_idx..end_idx]);
            start_idx += battery.0 + 1;
            battery.1
        })
        .collect();

    str::from_utf8(&batteries)
        .map_err(|e| e.to_string())?
        .parse::<usize>()
        .map_err(|e| e.to_string())
}

pub fn part1(input: &str) -> Result<usize, String> {
    banks(input).iter().map(|bank| joltage(bank, 2)).sum()
}

pub fn part2(input: &str) -> Result<usize, String> {
    banks(input).iter().map(|bank| joltage(bank, 12)).sum()
}

aocutils::assert_parts!(357, 16812, 3121910778619, 166345822896410);
