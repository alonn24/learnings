const MASS_TO_FUEL: fn(i32) -> i32 = |mass: i32| (mass / 3) - 2;

pub fn part1(input: Vec<i32>) -> i32 {
  input.iter().fold(0, |acc, mass| acc + MASS_TO_FUEL(*mass))
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_12() {
    assert_eq!(2, MASS_TO_FUEL(12));
  }

  #[test]
  fn calculate_100756() {
    assert_eq!(33583, MASS_TO_FUEL(100756));
  }
}
