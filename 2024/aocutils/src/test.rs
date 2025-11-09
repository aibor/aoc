#[macro_export]
macro_rules! assert_parts {
    ($p1_ex:expr, $p1_in:expr, $p2_ex:expr, $p2_in:expr) => {
        static EXAMPLE: &str = include_str!("../example");
        static INPUT: &str = include_str!("../input");

        #[test]
        fn part1_example() {
            assert_eq!($p1_ex, part1(EXAMPLE));
        }

        #[test]
        fn part1_input() {
            assert_eq!($p1_in, part1(INPUT));
        }

        #[test]
        fn part2_example() {
            assert_eq!($p2_ex, part2(EXAMPLE));
        }

        #[test]
        fn part2_input() {
            assert_eq!($p2_in, part2(INPUT));
        }
    };
}
