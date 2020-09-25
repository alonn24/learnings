const ADD: fn(usize, usize) -> usize = |x, y| x + y;
const MUL: fn(usize, usize) -> usize = |x, y| x * y;

fn get_swap_value(op_type: usize) -> fn(usize, usize) -> usize {
  match op_type {
    1 => ADD,
    2 => MUL,
    _ => panic!("Uh!"),
  }
}

fn get_output(mut input: Vec<usize>, noun: usize, verb: usize) -> Vec<usize> {
  input[1] = noun;
  input[2] = verb;
  let mut i = 0;
  loop {
    if i >= input.len() || input[i] == 99 {
      break;
    };
    let output_i = input[i + 3];
    input[output_i] = get_swap_value(input[i])(input[input[i + 1]], input[input[i + 2]]);
    i += 4;
  }
  input
}

pub fn part2(input: Vec<usize>) -> usize {
  for noun in 1..99 {
    for verb in 1..99 {
      let output = get_output(input.to_vec(), noun, verb);
      if output[0] == 19690720 {
        return noun * 100 + verb;
      }
    }
  }
  panic!("Could not found!");
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_12() {
    assert_eq!(vec![2, 0, 0, 0, 99], get_output(vec![1, 0, 0, 0, 99], 0, 0));
    assert_eq!(vec![2, 3, 0, 6, 99], get_output(vec![2, 0, 0, 3, 99], 3, 0));
  }
}
