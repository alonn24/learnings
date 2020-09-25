use std::fmt;

pub enum UserState {
  LoggedIn,
  LoggedOut,
}

pub struct User {
  name: String,
  user_state: UserState,
}

impl User {
  pub fn new(name: &str) -> User {
    User {
      name: String::from(name),
      user_state: UserState::LoggedOut,
    }
  }

  pub fn login(&mut self) {
    self.user_state = UserState::LoggedIn;
  }
}

impl fmt::Display for User {
  fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
    let state = match self.user_state {
      UserState::LoggedIn => "logged in",
      UserState::LoggedOut => "logged out",
    };
    write!(f, "name: {} is {}", self.name, state)
  }
}
