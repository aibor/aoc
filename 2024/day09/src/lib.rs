use std::{
    fmt::{Debug, Display},
    str::FromStr,
};

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
enum Block {
    File(usize),
    Free,
}

impl Block {
    fn is_free(&self) -> bool {
        matches!(self, Self::Free)
    }
}

impl Display for Block {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Self::File(id) => f.write_str(&(id % 10).to_string()),
            Self::Free => f.write_str("."),
        }
    }
}

struct BlockList(Vec<Block>);

impl Display for BlockList {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        for b in self.0.iter() {
            write!(f, "{b}")?
        }
        Ok(())
    }
}

impl BlockList {
    fn compact(&mut self) {
        let mut back = self.0.len() - 1;

        for front in 0..self.0.len() {
            if !self.0[front].is_free() {
                continue;
            }
            while self.0[back].is_free() {
                back -= 1;
            }
            if back < front {
                return;
            }
            self.0.swap(front, back);
            back -= 1;
        }
    }

    fn checksum(&self) -> usize {
        self.0.iter().enumerate().fold(0, |acc, (idx, b)| match b {
            Block::File(id) => acc + idx * *id,
            Block::Free => acc,
        })
    }
}

#[derive(Debug, Clone, Copy, PartialEq, Eq)]
struct File {
    offset: usize,
    len: usize,
    id: usize,
}

impl File {
    fn checksum(&self) -> usize {
        (0..self.len).fold(0, |acc, i| acc + (i + self.offset) * self.id)
    }

    fn next(&self) -> usize {
        self.offset + self.len
    }
}

#[derive(Debug, Clone, PartialEq, Eq)]
struct FS(Vec<File>);

impl FromStr for FS {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let mut fs = Vec::new();
        let mut offset = 0;
        let mut id = 0;

        for (idx, c) in s.char_indices().take_while(|(_, c)| c.is_numeric()) {
            let len = c.to_digit(10).ok_or("failed to parse len")? as usize;

            if idx % 2 == 0 {
                fs.push(File { offset, len, id });
                id += 1;
            };

            offset += len;
        }

        Ok(FS(fs))
    }
}

impl Display for FS {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        self.as_blocklist().fmt(f)
    }
}

impl FS {
    fn as_blocklist(&self) -> BlockList {
        let mut blocks = Vec::new();

        for file in self.0.iter() {
            let free = file.offset as isize - blocks.len() as isize;
            if free < 0 {
                panic!("offset below already written blocks");
            }
            if free > 0 {
                blocks.extend_from_slice(&[Block::Free].repeat(free as usize));
            }

            blocks.extend_from_slice(&[Block::File(file.id)].repeat(file.len));
        }

        BlockList(blocks)
    }

    fn find_space_left(&self, file: &File) -> Option<(usize, usize)> {
        let mut offset = 0;

        for (idx, f) in self.0.iter().enumerate() {
            if offset > f.offset {
                panic!("file offset too low for {file:?}, at {offset}");
            }
            if (f.offset - offset) >= file.len {
                return Some((idx, offset));
            };

            if file.id == f.id {
                break;
            }

            offset = f.next();
        }
        None
    }

    fn defrag(&mut self) {
        for file in self.0.clone().into_iter().rev() {
            let Some((idx, offset)) = self.find_space_left(&file) else {
                continue;
            };
            let Some((ridx, _)) = self.0.iter().enumerate().rfind(|f| f.1.id == file.id) else {
                continue;
            };
            let mut e = self.0.remove(ridx);
            e.offset = offset;
            self.0.insert(idx, e);
        }
    }

    fn checksum(&self) -> usize {
        self.0.iter().fold(0, |acc, b| acc + b.checksum())
    }
}

pub fn part1(input: &str) -> usize {
    let fs: FS = input.parse().expect("failed to parse input");
    let mut blocks = fs.as_blocklist();
    blocks.compact();
    blocks.checksum()
}

pub fn part2(input: &str) -> usize {
    let mut fs: FS = input.parse().expect("failed to parse input");
    fs.defrag();
    fs.checksum()
}

aocutils::assert_parts!(1928, 6279058075753, 2858, 6301361958738);
