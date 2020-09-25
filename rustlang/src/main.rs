// mod output;
// mod variables;
// mod flow_control;
// mod structures;
mod adventofcode2019;

#[cfg(test)]
#[macro_use]
extern crate pretty_assertions;

fn main() {
    // output::hello_world();
    // variables::basic();
    // flow_control::basic();
    // structures::basic();
    adventofcode2019::run();
}
