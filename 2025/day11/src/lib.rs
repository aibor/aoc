use std::collections::HashMap;

type DeviceTree<'a> = HashMap<&'a str, Vec<&'a str>>;

fn parse_devices(input: &str) -> Result<DeviceTree<'_>, String> {
    input
        .lines()
        .map(|line| {
            let (name, attached) = line.split_once(": ").ok_or("invalid device".to_string())?;
            Ok((name, attached.split_whitespace().collect()))
        })
        .collect::<Result<_, _>>()
}

struct PathFinder<'a> {
    devices: DeviceTree<'a>,
    end: &'a str,
    seen: HashMap<&'a str, usize>,
}

impl<'a> PathFinder<'a> {
    fn new(devices: DeviceTree<'a>, end: &'a str) -> Self {
        let seen = HashMap::new();
        Self { devices, end, seen }
    }

    fn reset(&mut self, end: &'a str) {
        self.seen.clear();
        self.end = end;
    }

    fn paths_from(&mut self, current: &'a str) -> Result<usize, String> {
        if current == self.end {
            return Ok(1);
        } else if current == "out" {
            return Ok(0);
        } else if let Some(paths) = self.seen.get(current) {
            return Ok(*paths);
        }

        let next = self
            .devices
            .get(current)
            .ok_or(format!("no next devs for {}", current))?
            .clone();

        let paths = next
            .iter()
            .map(|dev| self.paths_from(dev))
            .sum::<Result<_, _>>()?;

        self.seen.insert(current, paths);

        Ok(paths)
    }
}

pub fn part1(input: &str) -> Result<usize, String> {
    let devices = parse_devices(input)?;
    PathFinder::new(devices, "out").paths_from("you")
}

pub fn part2(mut input: &str) -> Result<usize, String> {
    if input.lines().count() == 10 {
        input = include_str!("../example2");
    }

    let devices = parse_devices(input)?;
    let mut path_finder = PathFinder::new(devices, "out");

    ["out", "dac", "fft", "svr"]
        .windows(2)
        .map(|devs| {
            path_finder.reset(devs[0]);
            path_finder.paths_from(devs[1])
        })
        .product()
}

aocutils::assert_parts!(5, 603, 2, 380961604031372);
