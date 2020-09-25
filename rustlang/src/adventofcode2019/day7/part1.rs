fn convert_input(input: String) -> Vec<i32> {
  input.split(",").map(|s| s.parse().unwrap()).collect()
}

fn create_combinations (options: &Vec<i32>) -> Vec<Vec<i32>> {
  let res = Vec::into_iter(*options).map(|x| {
    let without: Vec<i32> = Vec::into_iter(*options).filter(|&y| y != x).collect();
    let combinations = create_combinations(&without);
    let concatenated = [&vec![x], &without[..]].concat();
    concatenated
  });
  res.collect()
}

pub fn part1(raw_input: String) -> i32 {
  let input = convert_input(raw_input);
  let mut max = 0;
  let options = (0..5).collect::<Vec<i32>>();
  let combinations = create_combinations(&options);
  println!("{:?}", combinations);
  0
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test() {
    let input = "3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0".to_string();
    assert_eq!(part1(input), 43210);
  }
}
