extern crate regex;
use regex::Regex;
fn main() {
    let regex_string = Regex::new(r"(\d+)x(\d+)x(\d+)").unwrap();
    let stream_buffer = advent_of_code::file::read_stream();
    for line in stream_buffer.map(|line| {
        let defer_to_native_str = &*line.unwrap();
        let match_str = &regex_string.captures(defer_to_native_str).unwrap();
        (match_str[1].parse::<i32>().unwrap().to_owned(),
         match_str[2].parse::<i32>().unwrap().to_owned(),
         match_str[3].parse::<i32>().unwrap().to_owned()
        )
    }).map( |tuple_match|
        (2*tuple_match.0*tuple_match.1) + (2*tuple_match.1*tuple_match.2) + (2*tuple_match.0*tuple_match.2)
    ) {
        println!("{:?}", line);
    }
}