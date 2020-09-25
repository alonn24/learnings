use std::fs;
use std::iter::Iterator;
mod part1;
mod part2;

fn get_input() -> Vec<usize> {
  let contents = fs::read_to_string("./src/adventofcode2019/day2/day2.input")
    .expect("Something went wrong reading the file");
  contents.split(",").map(|s| s.parse().unwrap()).collect()
}

pub fn get_part1() -> usize {
  let mut input = get_input();
  input[1] = 12;
  input[2] = 2;
  let result = part1::part1(input);
  result[0]
}

pub fn get_part2() -> usize {
  let input = get_input();
  part2::part2(input)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(5305097, get_part1());
  }

  #[test]
  fn calculate_part2() {
    let input = get_input();
    assert_eq!(4925, part2::part2(input));
  }
}
