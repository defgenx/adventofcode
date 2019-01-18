extern crate regex;

use std::collections::HashMap;

use regex::Regex;

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let reader = advent_of_code::file::read_stream();
    let regex_string = Regex::new(r"^(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)$").unwrap();
    let mut grid: HashMap<(i32, i32), bool, _> = HashMap::with_capacity(1000 * 1000);
    for line in reader {
        let my_line = line.unwrap();
        let match_str = &regex_string.captures(&*my_line).unwrap();
        let first_dim = match_str[2].parse::<String>().unwrap();
        let sec_dim = match_str[3].parse::<String>().unwrap();
        let first_splited_dim: Vec<_> = first_dim.split(',').collect();
        let sec_splited_dim: Vec<_> = sec_dim.split(',').collect();
        let match_tuple = (match_str[1].parse::<String>().unwrap(), first_splited_dim, sec_splited_dim);

        let x = match_tuple.1.to_vec()[0].parse::<i32>().unwrap();
        let y = match_tuple.1.to_vec()[1].parse::<i32>().unwrap();
        let x1 = match_tuple.2.to_vec()[0].parse::<i32>().unwrap();
        let y1 = match_tuple.2.to_vec()[1].parse::<i32>().unwrap();
        match &*match_tuple.0 {
            "turn on" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        let my_tuple = (i, j);
                        grid.entry(my_tuple).or_insert(true);
                    }
                }
            }
            "turn off" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        if grid.contains_key(&(i, j)) {
                            grid.remove(&(i, j));
                        }
                    }
                }
            }
            "toggle" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        let my_tuple = (i, j);
                        if grid.contains_key(&my_tuple) {
                            grid.remove(&my_tuple);
                        } else {
                            grid.insert(my_tuple, true);
                        }
                    }
                }
            }
            _ => panic!("Tu bluffes Martoni !")
        }
    }
    println!("{:#?}", grid.len());
}

fn part_two() {
    let reader = advent_of_code::file::read_stream();
    let regex_string = Regex::new(r"^(turn on|turn off|toggle) (\d+,\d+) through (\d+,\d+)$").unwrap();
    let mut grid: HashMap<(i32, i32), i32, _> = HashMap::with_capacity(1000 * 1000);
    for line in reader {
        let my_line = line.unwrap();
        let match_str = &regex_string.captures(&*my_line).unwrap();
        let first_dim = match_str[2].parse::<String>().unwrap();
        let sec_dim = match_str[3].parse::<String>().unwrap();
        let first_splited_dim: Vec<_> = first_dim.split(',').collect();
        let sec_splited_dim: Vec<_> = sec_dim.split(',').collect();
        let match_tuple = (match_str[1].parse::<String>().unwrap(), first_splited_dim, sec_splited_dim);

        let x = match_tuple.1.to_vec()[0].parse::<i32>().unwrap();
        let y = match_tuple.1.to_vec()[1].parse::<i32>().unwrap();
        let x1 = match_tuple.2.to_vec()[0].parse::<i32>().unwrap();
        let y1 = match_tuple.2.to_vec()[1].parse::<i32>().unwrap();
        match &*match_tuple.0 {
            "turn on" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        let my_tuple = (i, j);
                        if grid.contains_key(&my_tuple) {
                            grid.insert(my_tuple, grid.get(&my_tuple).unwrap() + 1);
                        } else {
                            grid.insert(my_tuple, 1);
                        }
                    }
                }
            }
            "turn off" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        let my_tuple = (i, j);
                        if grid.contains_key(&my_tuple) && *grid.get(&my_tuple).unwrap() >= 1 {
                            grid.insert(my_tuple, grid.get(&my_tuple).unwrap() - 1);
                        }
                    }
                }
            }
            "toggle" => {
                for i in x..(x1 + 1) {
                    for j in y..(y1 + 1) {
                        let my_tuple = (i, j);
                        if grid.contains_key(&my_tuple) {
                            grid.insert(my_tuple, grid.get(&my_tuple).unwrap() + 2);
                        } else {
                            grid.insert(my_tuple, 2);
                        }
                    }
                }
            }
            _ => panic!("Tu bluffes Martoni !")
        }
    }
    let mut sum = 0;
    for val in grid.values() {
        sum += val;
    }
    println!("{:#?}", sum);
}