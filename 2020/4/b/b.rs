use std::collections::HashMap;
use regex::Regex;

fn main() {
    b();
}

fn b() {
    let mandatory_keys: Vec<&str> = vec!["byr",
                                         "iyr",
                                         "eyr",
                                         "hgt",
                                         "hcl",
                                         "ecl",
                                         "pid"
    ];

    let list = advent_of_code::file::read_stream();
    let mut store_passport: HashMap<String, String> = HashMap::new();
    let mut count_valid_passport = 0;
    for line in list {
        let unwrapped_line = line.unwrap();
        if unwrapped_line == "" {
            if is_valid_passport(store_passport, mandatory_keys.clone()) {
                count_valid_passport += 1;
            }
            store_passport = HashMap::new();
            continue;
        }
        let elts = unwrapped_line.split_whitespace();
        for elt in elts {
            let split: Vec<&str> = elt.split(":").collect();
            store_passport.insert(split[0].to_string(), split[1].to_string());
        }
    }
    // Handle EOF...
    if is_valid_passport(store_passport, mandatory_keys.clone()) {
        count_valid_passport += 1;
    }
    println!("The value is: {:#?}", count_valid_passport);
}

fn is_valid_passport(passport: HashMap<String, String>, mandatory_keys: Vec<&str>) -> bool {
    if !passport_mandatory_present(passport.clone(), mandatory_keys) {
        return false
    }
    for (key, val) in passport.iter() {
        if !passport_key_content_valid(key.to_string(), val.to_string()) {
            return false
        }
    }
    return true
}

fn passport_mandatory_present(passport: HashMap<String, String>, mandatory_keys: Vec<&str>) -> bool {
    let mut counter = 0;
    for key in &mandatory_keys {
        let vector = passport.keys().cloned().collect::<Vec<String>>();
        if vector.contains(&key.to_string()) {
            counter += 1
        }
    }
    if counter == mandatory_keys.len() {
        return true
    }
    return false
}

fn passport_key_content_valid(key: String, val: String) -> bool {
    if key == String::from("byr") {
        let int_val = val.parse::<i32>().unwrap();
        if int_val >= 1920 && int_val <= 2002 {
            return true
        }
    }
    else if key == String::from("iyr") {
        let int_val = val.parse::<i32>().unwrap();
        if int_val >= 2010 && int_val <= 2020 {
            return true
        }
    }
    else if key == String::from("eyr") {
        let int_val = val.parse::<i32>().unwrap();
        if int_val >= 2020 && int_val <= 2030 {
            return true
        }
    }
    else if key == String::from("hgt") {
        let re_cm = Regex::new(r"^(.+)cm$").unwrap();
        if re_cm.is_match(&val) {
            let capture = re_cm.captures(&val).unwrap();
            let length = capture[1].parse::<i32>().unwrap();
            if length >= 150 && length <= 193 {
                return true
            }
        }
        let re_in= Regex::new(r"(.+)in$").unwrap();
        if re_in.is_match(&val) {
            let capture = re_in.captures(&val).unwrap();
            let length = capture[1].parse::<i32>().unwrap();
            if length >= 59 && length <= 76 {
                return true
            }
        }
    }
    else if key == String::from("hcl") {
        let re = Regex::new(r"^#[0-9a-f]{6}$").unwrap();
        if re.is_match(&val) {
            return true
        }
    }
    else if key == String::from("ecl") {
        let colors: Vec<&str> = vec![
            "amb", "blu", "brn", "gry", "grn", "hzl", "oth",
        ];
        if colors.contains(&val.as_str()) {
            return true
        }
    }
    else if key == String::from("pid") {
        if val.chars().count() == 9 {
            return true
        }
    }
    else if key == String::from("cid") {
        return true
    }
    else {
        return false
    }
    return false
}