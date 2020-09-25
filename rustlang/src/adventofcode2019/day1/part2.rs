const MASS_TO_FUEL: fn(i32) -> i32 = |mass: i32| (mass / 3) - 2;

fn calculate_fuel_recursive(mass: i32) -> i32 {
  let fuel = MASS_TO_FUEL(mass);
  if fuel <= 0 { 0 }  else { fuel + calculate_fuel_recursive(fuel) }
}

// NOT 5040745, 5043715, 15129103
pub fn part2(input: Vec<i32>) -> i32 {
  input
    .iter()
    .fold(0, |acc, mass| acc + calculate_fuel_recursive(*mass))
}

#[cfg(test)]
mod tests {
  use super::*;

  #[test]
  fn calculate_fuel_of_12() {
    assert_eq!(2, calculate_fuel_recursive(12));
  }

  #[test]
  fn calculate_fuel_of_100756() {
    assert_eq!(50346, calculate_fuel_recursive(100756));
  }
}