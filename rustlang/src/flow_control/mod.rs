pub fn basic() {
  println!("~~~~~~~ Flow Control ~~~~~~~~~");
  let authenticated = false;
  let is_allowed = if authenticated { true } else { false };
  println!("Member allowed: {}", is_allowed);

  let add_nums = |n1: i32, n2: i32| n1 + n2;
  println!("1 + 2 = {}", add(1, 2));
  println!("1 + 2 = {}", add_nums(1, 2));

  let vec1 = vec![1, 2, 3];
  // When vec is borrowed we need to use it by reference
  let vec2 = &vec1;
  println!("borrowed values {:?}", (&vec1, vec2));
  println!("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
}

fn add(n1: i32, n2: i32) -> i32 {
  n1 + n2
}
