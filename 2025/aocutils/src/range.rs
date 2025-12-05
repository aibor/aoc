use std::ops::Range;

pub enum Overlap {
    None,
    Covered,
    Covers,
    Front,
    End,
}

pub fn overlap<T>(a: &Range<T>, b: &Range<T>) -> Overlap
where
    T: PartialOrd,
{
    if a.start >= b.start {
        if a.end <= b.end {
            Overlap::Covered
        } else if b.contains(&a.start) {
            Overlap::Front
        } else {
            Overlap::None
        }
    } else if a.end >= b.end {
        Overlap::Covers
    } else if b.contains(&a.end) {
        Overlap::End
    } else {
        Overlap::None
    }
}

pub struct MergedRangeIterator<T> {
    ranges: Vec<Range<T>>,
    index: usize,
}

impl<T> MergedRangeIterator<T>
where
    Self: Sized,
    T: Ord + Copy + Clone,
{
    fn new(mut ranges: Vec<Range<T>>) -> Self {
        // Sort so the merging algorithm needs to work forward only.
        ranges.sort_by(|a, b| a.start.cmp(&b.start));

        Self { ranges, index: 0 }
    }

    fn merge(&mut self) -> bool {
        if self.ranges[self.index].is_empty() {
            return false;
        }

        for other in (self.index + 1)..(self.ranges.len()) {
            let left = &self.ranges[self.index];
            let right = &self.ranges[other];

            match overlap(left, right) {
                Overlap::Covered => {
                    self.ranges[self.index].end = left.start;
                    return false;
                }
                Overlap::Covers => (),
                Overlap::Front => self.ranges[self.index].start = right.start,
                Overlap::End => self.ranges[self.index].end = right.end,
                _ => continue,
            }

            // Squash other when this range was extended or covered it already.
            self.ranges[other].end = self.ranges[other].start;
        }

        true
    }
}

impl<T> Iterator for MergedRangeIterator<T>
where
    Self: Sized,
    T: Ord + Copy + Clone,
{
    type Item = Range<T>;

    fn next(&mut self) -> Option<Self::Item> {
        loop {
            if self.index >= self.ranges.len() {
                break None;
            }

            let item = self.merge().then(|| self.ranges[self.index].clone());
            self.index += 1;

            if item.is_some() {
                break item;
            }
        }
    }
}

impl<T> FromIterator<Range<T>> for MergedRangeIterator<T>
where
    Self: Sized,
    T: Ord + Eq + Copy + Clone,
{
    fn from_iter<I: IntoIterator<Item = Range<T>>>(iter: I) -> Self {
        Self::new(iter.into_iter().collect())
    }
}

impl<T> From<&[Range<T>]> for MergedRangeIterator<T>
where
    Self: Sized,
    T: Ord + Copy + Clone,
{
    fn from(value: &[Range<T>]) -> Self {
        Self::new(value.to_vec())
    }
}

pub fn merged_ranges<T>(ranges: &[Range<T>]) -> MergedRangeIterator<T>
where
    T: Ord + Clone + Copy,
{
    MergedRangeIterator::from(ranges)
}
