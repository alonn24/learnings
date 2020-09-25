use std::fs;
use std::iter::Iterator;

mod part1;
mod part2;

fn get_input() -> Vec<i32> {
  let contents = fs::read_to_string("./src/adventofcode2019/day1/day1.input")
    .expect("Something went wrong reading the file");
  contents.split("\n").map(|s| s.parse().unwrap()).collect()
}

pub fn get_part1() -> i32 {
  let input = get_input();
  part1::part1(input)
}

pub fn get_part2() -> i32 {
  let input = get_input();
  part2::part2(input)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(3362507, part1::part1(get_input()));
  }

  #[test]
  fn calculate_part2() {
    assert_eq!(5040874, part2::part2(get_input()));
  }
}
