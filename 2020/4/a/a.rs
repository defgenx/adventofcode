fn main() {
    a();
}

fn a() {
    let mandatory_keys: Vec<&str> = vec!["byr",
                                            "iyr",
                                            "eyr",
                                            "hgt",
                                            "hcl",
                                            "ecl",
                                            "pid"
    ];

    let list = advent_of_code::file::read_stream();
    let mut store_passport: Vec<String> = Vec::new();
    let mut count_valid_passport = 0;
    for line in list {
        let unwrapped_line = line.unwrap();
        if unwrapped_line == "" {
            if is_valid_passport(store_passport, mandatory_keys.clone()) {
                count_valid_passport += 1;
            }
            store_passport = Vec::new();
            continue;
        }
        let elts = unwrapped_line.split_whitespace();
        for elt in elts {
            let split: Vec<&str> = elt.split(":").collect();
            store_passport.push(split[0].to_string())
        }
    }
    if is_valid_passport(store_passport, mandatory_keys.clone()) {
        count_valid_passport += 1;
    }
    println!("The value is: {:#?}", count_valid_passport);
}

fn is_valid_passport(keys: Vec<String>, mandatory_keys: Vec<&str>) -> bool {
    let mut counter = 0;
    for key in &mandatory_keys {
        if keys.contains(&key.to_string()) {
            counter += 1
        }
    }
    if counter == mandatory_keys.len() {
        return true
    }
    return false
}