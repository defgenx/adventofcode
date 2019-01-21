extern crate regex;

use regex::Regex;

fn main() {
    part_one();
}

fn part_one() {
    let reader = advent_of_code::file::read_stream();
    let regex_string_not = Regex::new(r"^NOT ([a-z0-9]+)$").unwrap();
    let regex_string_others = Regex::new(r"^(.+) (LSHIFT|RSHIFT|OR|AND) (.+)$").unwrap();
    let mut data = Vec::new();
    for line in reader {;
        let mut local_data = Vec::new();
        let my_line = line.unwrap();
        let split_str: Vec<_> = my_line.split(" -> ").collect();
        let left_str = split_str[0];
        let right_str = split_str[1];
        let match_str;
        if left_str.contains("NOT") {
            match_str = regex_string_not.captures(&*left_str).unwrap();
            let captured: String = match_str[1].parse().unwrap();
            local_data.push(captured);
        } else if left_str.contains("LSHIFT")
            || left_str.contains("RSHIFT")
            || left_str.contains("OR")
            || left_str.contains("AND")
        {
            match_str = regex_string_others.captures(&*split_str[0]).unwrap();
            local_data.push(match_str[1].parse().unwrap());
            local_data.push(match_str[3].parse().unwrap());
        } else {
            local_data.push(left_str.to_string());
        }
        local_data.push(right_str.to_string());
        data.push(local_data);
        println!("{:#?}", data);
    }
}

fn is_numeric(captured: String) -> bool {
    captured.parse::<i32>().is_ok()
}