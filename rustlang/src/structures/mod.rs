pub mod user;
use user::User;

pub fn basic() {
  println!("~~~~~~~ Structures ~~~~~~~~~");
  let mut user = User::new("Alon");
  println!("User {}", user);
  user.login();
  println!("User {}", user);
  println!("~~~~~~~~~~~~~~~~~~~~~~~~~~~~");
}
