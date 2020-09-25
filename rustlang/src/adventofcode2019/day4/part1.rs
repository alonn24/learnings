fn is_valid(input: String) -> bool {
  let mut streak = 0;
  let mut double = false;
  input
    .chars()
    .map(|x| x.to_digit(10).unwrap())
    .fold(9, |prev, cur| {
      double = prev == cur || double;
      if prev <= cur {
        streak += 1;
      } else {
        streak = 0;
        double = false;
      }
      cur
    });
  streak == 5 && double
}

// <1729
pub fn part1(start: i32, end: i32) -> i32 {
  let mut  count = 0;
  for num in start..=end {
    if is_valid(num.to_string()) {
      count += 1;
    }
  }
  count
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_simple() {
    assert_eq!(1, part1(111111, 111111));
    // 55, 56, 57, 56, 59
    assert_eq!(5, part1(113450, 113460));
    assert_eq!(0, part1(223450, 223450));
    assert_eq!(0, part1(123789, 123789));
  }
}
