fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let mut trees = 0;
    let mut right = 3;
    let vec_rows: Vec<String> =
        list.map(|line| {
            line.unwrap()
        }).collect();
    for (index, row) in vec_rows.iter().enumerate() {
        if index == 0 {
            continue;
        }

        if (row.chars().count() - 1) < right {
            right = right - row.chars().count()
        }

        if row.chars().nth(right).unwrap().to_string() == "#" {
            trees += 1
        }
        right += 3
    }
    println!("The value is: {:#?}", trees);
}