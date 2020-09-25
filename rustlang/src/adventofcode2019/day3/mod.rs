use std::fs;
use std::iter::Iterator;
mod part1;
mod part2;

fn get_input() -> (String, String) {
  let contents = fs::read_to_string("./src/adventofcode2019/day3/day3.input")
    .expect("Something went wrong reading the file");
  let wires: Vec<String> = contents.split("\n").map(|s| s.parse().unwrap()).collect();
  (
    wires[0].parse().unwrap(),
    wires[1].parse().unwrap()
  )
}

pub fn get_part1() -> i32 {
  let wires = get_input();
  part1::part1(wires.0, wires.1)
}

pub fn get_part2() -> i32 {
  let wires = get_input();
  part2::part2(wires.0, wires.1)
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_part1() {
    assert_eq!(896, get_part1());
  }

  #[test]
  fn calculate_part2() {
    assert_eq!(16524, get_part2());
  }
}
