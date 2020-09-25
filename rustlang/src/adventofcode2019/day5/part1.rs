const ADD: fn(i32, i32) -> i32 = |x, y| x + y;
const MUL: fn(i32, i32) -> i32 = |x, y| x * y;

fn get_swap_value(op_type: i32) -> fn(i32, i32) -> i32 {
  match op_type {
    1 => ADD,
    2 => MUL,
    _ => panic!("Uh!"),
  }
}

const BASE: i32 = 10;
fn get_value_by_type(values: &Vec<i32>, i: i32, offset: i32) -> i32 {
  let value = values[(i + offset) as usize];
  let op_code = values[i as usize];
  // save two first spots to ops
  let ref_type = (op_code % BASE.pow(offset as u32 + 2)) / BASE.pow(offset as u32 + 1);
  return if ref_type == 1 { value } else { values[value as usize] };
}

pub fn part1(input: Vec<i32>, op_input: i32) -> i32 {
  let mut result = input.to_vec();
  let mut i = 0;
  let mut output = 0;
  loop {
    if i >= result.len() || result[i] == 99 {
      break;
    };
    let op = result[i];
    let op_type = op % 100;
    if op_type == 3 || op_type == 4 {
      let ref_index = result[i + 1];
      if op_type == 3 {
        result[ref_index as usize] = op_input;
      } else {
        output = result[ref_index as usize];
      }
      i += 2;
    } else {
      let output_i = result[i + 3];
      let first_value = get_value_by_type(&result, i as i32, 1);
      let second_value = get_value_by_type(&result, i as i32, 2);
      result[output_i as usize] = get_swap_value(op_type)(first_value, second_value);
      i += 4;
    }
  }
  output
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn test_get_value_by_type() {
    assert_eq!(get_value_by_type(&vec![0, 2, 99], 0, 1), 99);
    assert_eq!(get_value_by_type(&vec![101, 2, 99], 0, 1), 2);
  }

  #[test]
  fn simple_test() {
    assert_eq!(part1(vec![4, 0, 99], 1), 4);
    assert_eq!(part1(vec![3, 0, 4, 0, 99], 1), 1);
    assert_eq!(part1(vec![3, 0, 1, 0, 0, 0, 4, 0, 99], 1), 2);
  }
}
