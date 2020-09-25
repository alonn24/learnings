pub fn hello_world() {
  println!("~~~~~~~ Output ~~~~~~~~~");
  println!("Hello {who}!", who = "World");
  println!("{:?}", (std::i32::MAX, std::i32::MIN));
  println!("~~~~~~~~~~~~~~~~~~~~~~~~\n");
}