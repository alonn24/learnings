const ADD: fn(usize, usize) -> usize = |x, y| x + y;
const MUL: fn(usize, usize) -> usize = |x, y| x * y;

fn get_swap_value(op_type: usize) -> fn(usize, usize) -> usize {
  match op_type {
    1 => ADD,
    2 => MUL,
    _ => panic!("Uh!"),
  }
}
// NOT 103, 89
pub fn part1(input: Vec<usize>) -> Vec<usize> {
  let mut result = input.to_vec();
  let mut i = 0;
  loop {
    if i >= result.len() || result[i] == 99 {
      break;
    };
    let output_i = result[i + 3];
    result[output_i] = get_swap_value(result[i])(result[result[i + 1]], result[result[i + 2]]);
    i += 4;
  }
  result
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_12() {
    assert_eq!(vec![2, 0, 0, 0, 99], part1(vec![1, 0, 0, 0, 99]));
    assert_eq!(vec![2, 3, 0, 6, 99], part1(vec![2, 3, 0, 3, 99]));
  }
}
