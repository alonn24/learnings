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
  return if ref_type == 1 {
    value
  } else {
    values[value as usize]
  };
}

pub fn part2(input: &Vec<i32>, op_input: i32) -> i32 {
  let mut result = input.to_vec();
  let mut i = 0;
  let mut counter = 0;
  let mut output = 0;
  loop {
    if counter >= 1000 {
      panic!("too much loop");
    }
    counter += 1;

    let op = result[i];
    let op_type = op % 100;
    if result[i] == 99 {
      break;
    } else if op_type >= 1 && op_type <= 2 {
      let output_i = result[i + 3];
      let first_value = get_value_by_type(&result, i as i32, 1);
      let second_value = get_value_by_type(&result, i as i32, 2);
      result[output_i as usize] = get_swap_value(op_type)(first_value, second_value);
      i += 4;
    } else if op_type == 3 {
      let ref_index = result[i + 1];
      result[ref_index as usize] = op_input;
      i += 2;
    } else if op_type == 4 {
      output = get_value_by_type(&result, i as i32, 1);
      i += 2;
    } else if op_type >= 5 && op_type <= 6 {
      let value = get_value_by_type(&result, i as i32, 1);
      let jump_value = get_value_by_type(&result, i as i32, 2);
      let should_jump = (op_type == 5 && value != 0) || (op_type == 6 && value == 0);
      if should_jump {
        i = jump_value as usize;
      } else {
        i += 3;
      };
    } else if op_type >= 7 && op_type <= 8 {
      let value1 = get_value_by_type(&result, i as i32, 1);
      let value2 = get_value_by_type(&result, i as i32, 2);
      let store = op_type == 7 && value1 < value2 || op_type == 8 && value1 == value2;
      let output_i = result[i + 3];
      result[output_i as usize] = if store { 1 } else { 0 };
      i += 4;
    } else {
      panic!("what is this op type? {}:{}", i, op_type);
    }
  }
  output
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn simple_test() {
    assert_eq!(part2(&vec![3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8], 8), 1);
    assert_eq!(part2(&vec![3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8], 9), 0);
    assert_eq!(part2(&vec![3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8], 7), 1);
    assert_eq!(part2(&vec![3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8], 8), 0);
  }

  #[test]
  fn larger_test() {
    let input = &vec![
      3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0,
      1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105,
      1, 46, 98, 99,
    ];
    assert_eq!(part2(input, 7), 999);
    assert_eq!(part2(input, 8), 1000);
    assert_eq!(part2(input, 9), 1001);
  }
}
