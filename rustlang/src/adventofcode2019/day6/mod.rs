use std::fs;
mod part1;
mod part2;

fn get_input() -> String {
  fs::read_to_string("./src/adventofcode2019/day6/day6.input")
    .expect("Something went wrong reading the file")
}

pub fn get_part1() -> i32 {
  part1::part1(get_input())
}

pub fn get_part2() -> i32 {
  part2::part2(get_input())
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(253104, get_part1());
  }

  #[test]
  fn calculate_part2() {
    assert_eq!(499, get_part2());
  }
}
