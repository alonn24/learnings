pub fn basic() {
  println!("~~~~~~~ Variables ~~~~~~~~~");
  const ID: i32 = 001;

  let name = String::from("Alon");
  let person: (String, &str, i8) = (name, "Yehezkel", 32);
  println!(
    "My name is {name} {family} ID:{id}",
    name = person.0,
    family = person.1,
    id = ID
  );

  let mut age = 30;
  println!("I am {age}yo", age = age);
  age = 32;
  println!("Actually my age is really {age}", age = age);

  let numbers: [i32; 4] = [1, 2, 3, 4];
  let first_numbers: &[i32] = &numbers[0..2]; // explusive on the right side
  println!("The first numbers are {:?}", first_numbers);

  let mut vec: Vec<i32> = vec![1, 2, 3, 4];
  let mut sum = 0;
  for x in vec.iter_mut() {
    *x += 1;
  }
  for x in vec.iter() {
    sum += x;
  }
  println!("Sum in {sum}", sum = sum);
  println!("~~~~~~~~~~~~~~~~~~~~~~~~~~~\n");
}
