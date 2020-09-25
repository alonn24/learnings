use std::fs;
use std::iter::Iterator;
mod part1;
mod part2;

fn get_input() -> Vec<i32> {
  let contents = fs::read_to_string("./src/adventofcode2019/day5/day5.input")
    .expect("Something went wrong reading the file");
  contents.split(",").map(|s| s.parse().unwrap()).collect()
}

pub fn get_part1() -> i32 {
  part1::part1(get_input(), 1)
}

pub fn get_part2() -> i32 {
  part2::part2(&get_input(), 5)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(7988899, get_part1());
  }

  #[test]
  fn calculate_part2() {
    assert_eq!(13758663, get_part2());
  }
}
