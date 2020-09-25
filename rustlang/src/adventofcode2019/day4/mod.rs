mod part1;
mod part2;

pub fn get_part1() -> i32 {
  part1::part1(372304, 847060)
}

pub fn get_part2() -> i32 {
  part2::part2(372304, 847060)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(475, get_part1());
  }

  #[test]
  fn calculate_part2() {
    assert_eq!(297, get_part2());
  }
}