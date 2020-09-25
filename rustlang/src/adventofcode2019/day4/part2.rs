fn is_valid(input: String) -> bool {
  let mut streak = 0;
  let mut double_streak = 0;
  let mut double = false;
  input
    .chars()
    .map(|x| x.to_digit(10).unwrap())
    .fold(9, |prev, cur| {
      if prev == cur {
        double_streak += 1;
      } else  {
        double = double_streak == 1 || double;
        double_streak = 0;
      }
      
      if prev <= cur {
        streak += 1;
      } else {
        streak = 0;
        double = false;
      }
      cur
    });
  double = double_streak == 1 || double;
  streak == 5 && double
}

pub fn part2(start: i32, end: i32) -> i32 {
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
    assert_eq!(1, part2(112233, 112233));
    assert_eq!(0, part2(123444, 123444));
    assert_eq!(1, part2(111122, 111122));
  }
}
