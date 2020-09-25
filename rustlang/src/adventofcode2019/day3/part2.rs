use std::collections::HashMap;

fn parse_wire(wire: &String) -> Vec<(String, i32)> {
  wire
    .split(",")
    .map(|s| {
      let first = &s[..1];
      let second = &s[1..];
      (String::from(first), second.parse().unwrap())
    })
    .collect()
}

fn make_move(dir: &String, current: (i32, i32)) -> (i32, i32) {
  let (x, y) = current;
  return if dir == "U" {
    (x, y - 1)
  } else if dir == "D" {
    (x, y + 1)
  } else if dir == "L" {
    (x - 1, y)
  } else {
    (x + 1, y)
  };
}

pub fn part2(str_wire1: String, str_wire2: String) -> i32 {
  let wire1 = parse_wire(&str_wire1);
  let wire2 = parse_wire(&str_wire2);
  let mut visited: HashMap<(i32, i32), i32> = HashMap::new();
  let mut current = (0, 0);
  let mut count = 0;
  for step in wire1.iter() {
    let (d, n) = step;
    for _i in 0..*n {
      current = make_move(d, current);
      count += 1;
      if !visited.contains_key(&current) {
        visited.insert(current, count);
      }
    }
  }
  
  let mut found: Vec<((i32, i32))> = vec![];
  count = 0;
  current = (0, 0);
  for step in wire2.iter() {
    let (d, n) = step;
    for _i in 0..*n {
      current = make_move(d, current);
      count += 1;
      match visited.get(&current) {
        Some(value) => found.push((count, *value)),
        None => {}
      }
    }
  }
  found.iter().fold(std::i32::MAX, |acc, item| i32::min(acc, item.0 + item.1))
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_simple() {
    let wire1 = String::from("R8,U5,L5,D3");
    let wire2 = String::from("U7,R6,D4,L4");
    assert_eq!(30, part2(wire1, wire2));
  }

  #[test]
  fn calculate_example1() {
    let wire1 = String::from("R75,D30,R83,U83,L12,D49,R71,U7,L72");
    let wire2 = String::from("U62,R66,U55,R34,D71,R55,D58,R83");
    assert_eq!(610, part2(wire1, wire2));
  }

  #[test]
  fn calculate_example2() {
    let wire1 = String::from("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51");
    let wire2 = String::from("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7");
    assert_eq!(410, part2(wire1, wire2));
  }
}
