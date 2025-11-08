use std::{cmp::Ordering, collections::HashSet};

#[derive(Debug, Eq, Hash, PartialEq)]
struct OrderingRule<'a>(&'a str, &'a str);

impl<'a> OrderingRule<'a> {
    fn from(str: &str) -> Option<OrderingRule<'_>> {
        let mut s = str.split("|");
        Some(OrderingRule(s.next()?, s.next()?))
    }
}

#[derive(Debug, Clone)]
struct Pages<'a>(Vec<&'a str>);

impl<'a> Pages<'a> {
    fn mid_page(&self) -> usize {
        let mid_index = self.0.len() / 2;
        let mid_page = self.0.get(mid_index).unwrap();
        mid_page.parse::<usize>().unwrap()
    }

    fn sort_by(&mut self, compare: impl FnMut(&&str, &&str) -> Ordering) {
        self.0.sort_by(compare);
    }
}

struct OrderingRules<'a>(HashSet<OrderingRule<'a>>);

impl<'a> OrderingRules<'a> {
    fn is_sorted(&self, pages: &Pages) -> bool {
        pages
            .0
            .windows(2)
            .all(|w| self.cmp_pages(w[0], w[1]) == Ordering::Less)
    }

    fn cmp_pages(&self, a: &str, b: &str) -> Ordering {
        if self.0.contains(&OrderingRule(a, b)) {
            Ordering::Less
        } else if self.0.contains(&OrderingRule(b, a)) {
            Ordering::Greater
        } else {
            Ordering::Equal
        }
    }
}

type Updates<'a> = Vec<Pages<'a>>;

fn parse_input(input: &str) -> (OrderingRules<'_>, Updates<'_>) {
    let mut lines = input.lines();

    let rules = HashSet::from_iter(
        (&mut lines)
            .take_while(|line| !line.is_empty())
            .filter_map(OrderingRule::from),
    );

    let updates = lines.map(|line| Pages(line.split(",").collect())).collect();

    (OrderingRules(rules), updates)
}

pub fn part1(input: &str) -> usize {
    let (rules, updates) = parse_input(input);

    updates
        .iter()
        .filter(|update| rules.is_sorted(update))
        .map(Pages::mid_page)
        .sum()
}

pub fn part2(input: &str) -> usize {
    let (rules, updates) = parse_input(input);

    updates
        .into_iter()
        .filter(|update| !rules.is_sorted(update))
        .map(|mut pages| {
            pages.sort_by(|a, b| rules.cmp_pages(a, b));
            pages.mid_page()
        })
        .sum()
}

aocutils::assert_parts!(143, 5762, 123, 4130);
