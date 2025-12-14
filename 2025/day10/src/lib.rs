use std::{fmt::Display, str::FromStr};

use itertools::Itertools;

#[derive(Debug, Clone)]
struct Machine {
    components: usize,
    indicator_lights: u16,
    buttons: Vec<u16>,
    joltages: Vec<u16>,
}

impl Display for Machine {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "[{0:01$b}] {2:?} {{{3:?}}}",
            self.indicator_lights, self.components, self.buttons, self.joltages
        )
    }
}

impl FromStr for Machine {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut parts = s.split_whitespace();
        let indicator_lights = parts.next().ok_or("no indicator lights".to_string())?;
        let joltages = parts.next_back().ok_or("no joltages".to_string())?;
        let buttons: Vec<&str> = parts.collect();

        let components = indicator_lights.len() - 2;
        let indicator_lights = indicator_lights[1..(indicator_lights.len() - 1)]
            .chars()
            .enumerate()
            .fold(0, |acc, (idx, c)| match c {
                '#' => acc | (1 << idx),
                _ => acc,
            });

        let buttons = buttons
            .iter()
            .map(|s| {
                s[1..(s.len() - 1)].split(",").try_fold(0, |button, s| {
                    let n = s.parse::<u8>().map_err(|e| e.to_string() + s)?;
                    Ok(button | (1 << n))
                })
            })
            .collect::<Result<Vec<u16>, String>>()?;

        let joltages = joltages[1..(joltages.len() - 1)]
            .split(",")
            .map(|s| s.parse::<u16>().map_err(|e| e.to_string() + s))
            .collect::<Result<Vec<u16>, String>>()?;

        Ok(Self {
            components,
            indicator_lights,
            buttons,
            joltages,
        })
    }
}

impl Machine {
    fn find_presses(&self, pattern: &u16) -> Option<Vec<u16>> {
        for i in 1..self.buttons.len() {
            let mut solutions = Vec::new();
            for buttons in self.buttons.iter().combinations(i) {
                let mut state = 0;
                for (idx, button) in buttons.iter().enumerate() {
                    state ^= *button;
                    if &state != pattern {
                        continue;
                    }
                    solutions.push(Vec::from_iter(buttons[..=idx].iter().map(|&&b| b)));
                }
            }
            if !solutions.is_empty() {
                return solutions.into_iter().min_by_key(|b| b.len());
            }
        }
        None
    }
}

fn parse_machines(input: &str) -> Result<Vec<Machine>, String> {
    input.lines().map(Machine::from_str).collect()
}

pub fn part1(input: &str) -> Result<usize, String> {
    let machines = parse_machines(input)?;

    let presses = machines
        .iter()
        .filter_map(|m| m.find_presses(&m.indicator_lights).map(|b| b.len()))
        .sum();

    Ok(presses)
}

pub fn part2(input: &str) -> Result<usize, String> {
    _ = input;
    Ok(0)
}

aocutils::assert_parts!(7, 514, 0, 0);
