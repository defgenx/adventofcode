use std::collections::HashMap;

fn main() {
    b();
}

fn b() {
    let list = advent_of_code::file::read_stream();
    let mut map_rows: Vec<HashMap<&str, String>> = Vec::new();
    for line in list {
        let mut map_row = HashMap::new();
        let split_line = line.as_ref().unwrap().split_whitespace();
        let elts: Vec<&str> = split_line.collect();
        let min_max: Vec<&str> = elts[0].split("-").collect();
        let letters: Vec<&str> = elts[1].split(":").collect();
        map_row.insert(
            "min",
            min_max[0].to_string()
        );
        map_row.insert(
            "max",
            min_max[1].to_string()
        );
        map_row.insert(
            "letter",
            letters[0].to_string()
        );
        map_row.insert(
            "pwd",
            elts[2].to_string()
        );
        map_rows.push(map_row)
    };
    let mut total_valid_pwd: i32 = 0;
    for row in map_rows {
        let min: usize = row["min"].parse::<usize>().unwrap();
        let max: usize = row["max"].parse::<usize>().unwrap();
        let min_letter: String = row["pwd"].chars().nth(min -1).unwrap().to_string();
        let max_letter: String = row["pwd"].chars().nth(max -1).unwrap().to_string();
        if min_letter == row["letter"] && max_letter == row["letter"] {
            continue;
        }
        if min_letter == row["letter"] || max_letter == row["letter"] {
            println!("{:#?}", row);
            total_valid_pwd += 1;
        }
    }
    println!("The value is: {:#?}", total_valid_pwd);
}