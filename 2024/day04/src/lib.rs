use std::collections::{HashMap, HashSet};

struct WordSearch {
    chars: HashMap<char, HashSet<isize>>,
    line: isize,
}

impl WordSearch {
    fn new(input: &str) -> WordSearch {
        let mut chars = HashMap::new();

        for (idx, c) in input.char_indices() {
            chars
                .entry(c)
                .or_insert(HashSet::new())
                .insert(idx as isize);
        }

        let line_len = input.find('\n').expect("no linebreak found");
        let line = (line_len as isize) + 1;

        WordSearch { chars, line }
    }

    fn positions(&self, c: char) -> &HashSet<isize> {
        self.chars.get(&c).unwrap()
    }

    fn is_at(&self, c: char, pos: isize) -> bool {
        self.chars[&c].contains(&pos)
    }
}

fn adjacents_dists(abs_dists: &[isize]) -> Vec<isize> {
    Vec::from_iter(abs_dists.iter().flat_map(|a| [*a, -a]))
}

pub fn part1(input: &str) -> usize {
    let ws = WordSearch::new(input);

    // All adjacent letter distances. Positives defined, negatives calculated.
    let adj = adjacents_dists(&[1, ws.line - 1, ws.line, ws.line + 1]);

    let is_xmas = |pos_x: &isize, dist: &isize| {
        ws.is_at('M', pos_x + dist)
            && ws.is_at('A', pos_x + 2 * dist)
            && ws.is_at('S', pos_x + 3 * dist)
    };

    let count_xmas_from = |pos_x: &isize| adj.iter().filter(|dist| is_xmas(pos_x, dist)).count();

    // Find an X and then try finding a MAS in each direction.
    ws.positions('X').iter().map(count_xmas_from).sum()
}

pub fn part2(input: &str) -> usize {
    let ws = WordSearch::new(input);

    // Diagonals only. Positives defined, negatives calculated.
    let adj = adjacents_dists(&[ws.line - 1, ws.line + 1]);

    let is_mas = |pos_a: &isize, dist: &isize| -> bool {
        ws.is_at('M', pos_a + dist) && ws.is_at('S', pos_a - dist)
    };

    let has_x = |pos_a: &isize, first_dist: &isize| {
        adj.iter()
            .filter(|other_dist| other_dist.abs() != first_dist.abs())
            .any(|dist| is_mas(pos_a, dist))
    };

    let is_x_mas = |pos_a, dist| is_mas(pos_a, dist) && has_x(pos_a, dist);

    ws.positions('A')
        .iter()
        .filter(|pos_a| adj.iter().any(|dist| is_x_mas(pos_a, dist)))
        .count()
}

aocutils::assert_parts!(18, 2358, 9, 1737);
