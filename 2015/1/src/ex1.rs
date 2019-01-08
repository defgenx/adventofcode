use std::io::prelude::*;
use std::path::Path;
use std::fs::File;
use std::error::Error;
use adventofcode::*;

fn main() {


    // Create a path
    let path = Path::new("input_ex1.txt");
    let display = path.display();

    // Open file in ro mode
    let mut file = match File::open(&path) {
        Err(why) => panic!("couldn't open exercise file {}: {}", display,
                           why.description()),
        Ok(file) => file,
    };

    // Init string var
    let mut s = String::new();
    match file.read_to_string(&mut s) {
        Err(why) => panic!("couldn't read {}: {}", display,
                           why.description()),
        Ok(_) => print!("{} contains:\n{}", display, s),
    }

    print!("{:?}", file);
}