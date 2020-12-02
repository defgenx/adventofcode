use std::collections::HashMap;

fn main() {
    a();
}

fn a() {
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
        let min: i32 = row["min"].parse::<i32>().unwrap();
        let max: i32 = row["max"].parse::<i32>().unwrap();
        let count: i32 = row["pwd"].matches(row["letter"].as_str()).count() as i32;
        if count >= min && count <= max {
            println!("{:#?}", row);
            total_valid_pwd += 1;
        }
    }
    println!("The value is: {:#?}", total_valid_pwd);
}