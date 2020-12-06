use std::collections::HashSet;

fn main() {
    b();
}

fn b() {
    let list = advent_of_code::file::read_stream();
    let mut hash_by_group: HashSet<String> = HashSet::new();
    let mut total: usize = 0;
    let mut first_iter = true;
    for line in list {
        let unwrapped_line = line.unwrap();
        if unwrapped_line == "" {
            hash_by_group = HashSet::new();
            first_iter = true;
            continue;
        }
        let elts: Vec<char> = unwrapped_line.chars().collect();
        let mut tmp_hash_by_group = HashSet::new();
        for elt in elts {
            if first_iter {
                hash_by_group.insert(elt.to_string());
                total += 1
            }
            tmp_hash_by_group.insert(elt.to_string());
        }
        let hash_len = hash_by_group.len();
        hash_by_group = intersection(hash_by_group.clone(), &tmp_hash_by_group);
        total -= hash_len - hash_by_group.len();
        first_iter = false
    }
    println!("Value is: {:#?}", total);
}

// Useless to use generic here
fn intersection(a: HashSet<String>, b: &HashSet<String>) -> HashSet<String> {
    a.into_iter().filter(|e| b.contains(e)).collect()
}