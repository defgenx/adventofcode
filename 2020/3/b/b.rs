use std::collections::HashMap;

fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let mut total_trees: i64 = 1;
    let vec_rows: Vec<String> =
        list.map(|line| {
            line.unwrap()
        }).collect();
    let mut slopes: Vec<HashMap<&str, i32>> = Vec::new();
    slopes.push(
        [("right", 1),
            ("down", 1)]
            .iter().cloned().collect()
    );
    slopes.push(
        [("right", 3),
            ("down", 1)]
            .iter().cloned().collect()
    );
    slopes.push(
        [("right", 5),
            ("down", 1)]
            .iter().cloned().collect()
    );
    slopes.push(
        [("right", 7),
            ("down", 1)]
            .iter().cloned().collect()
    );

    slopes.push(
        [("right", 1),
            ("down", 2)]
            .iter().cloned().collect()
    );

    for slope in slopes {
        let mut right: usize = slope["right"] as usize;
        let mut trees: i64 = 0;
        for (index, row) in vec_rows.iter().enumerate() {
            if index % slope["down"] as usize != 0 || index == 0 {
                continue;
            }

            if (row.chars().count() - 1) < right {
                right = right - row.chars().count();
            }

            if row.chars().nth(right).unwrap().to_string() == "#" {
                trees += 1;
            }
            right += 3;
        }
        println!("{:#?}", slope);
        total_trees *= trees;
    }
    println!("The value is: {:#?}", total_trees);
}