use std::collections::HashSet;

fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let mut hash_by_group: HashSet<String> = HashSet::new();
    let mut total = 0;
    for line in list {
        let unwrapped_line = line.unwrap();
        if unwrapped_line == "" {
            hash_by_group = HashSet::new();
            continue;
        }
        let elts: Vec<char> = unwrapped_line.chars().collect();
        for elt in elts {
            if !hash_by_group.contains(&elt.to_string()) {
                total += 1;
            }
            hash_by_group.insert(elt.to_string());

        }
    }
    println!("Value is: {:#?}", total);
}