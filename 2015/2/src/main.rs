extern crate regex;
use regex::Regex;
use std::collections::HashMap;
fn main() {
    let regex_string = Regex::new(r"(\d+)x(\d+)x(\d+)").unwrap();
    let stream_buffer = advent_of_code::file::read_stream();
    let mut sum = 0;
    for line in stream_buffer.map(|line| {
        let defer_to_native_str = &*line.unwrap();
        let match_str = &regex_string.captures(defer_to_native_str).unwrap();
        (match_str[1].parse::<i32>().unwrap().to_owned(),
         match_str[2].parse::<i32>().unwrap().to_owned(),
         match_str[3].parse::<i32>().unwrap().to_owned()
        )
    }).map( |tuple_match| {
        let first = tuple_match.0 * tuple_match.1;
        let second = tuple_match.1 * tuple_match.2;
        let third = tuple_match.0 * tuple_match.2;
        let array_tmp = [first, second, third];
        let min = array_tmp.iter().min();
        (2 * tuple_match.0 * tuple_match.1) + (2 * tuple_match.1 * tuple_match.2) + (2 * tuple_match.0 * tuple_match.2) + min.unwrap()
    }) {
        sum += line;
    }
    println!("{:?}", sum);


    let stream_buffer_part_two = advent_of_code::file::read_stream();
    let mut sum_ribbon = 0;
    for line in stream_buffer_part_two.map(|line| {
        let defer_to_native_str = &*line.unwrap();
        let match_str = &regex_string.captures(defer_to_native_str).unwrap();
        [match_str[1].parse::<i32>().unwrap().to_owned(),
         match_str[2].parse::<i32>().unwrap().to_owned(),
         match_str[3].parse::<i32>().unwrap().to_owned()
        ]
    }).map( |array_match| {
        let mut sorted_array = array_match;
        sorted_array.sort();
        (2 * sorted_array[0]) + (2 * sorted_array[1]) + (sorted_array[0] * sorted_array[1] * sorted_array[2])
    }) {
        sum_ribbon += line;
    }
    println!("{:?}", sum_ribbon);

}