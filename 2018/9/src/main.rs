extern crate regex;
use regex::Regex;
use std::collections::VecDeque;

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let regex_string = Regex::new(r"^(\d+) players; last marble is worth (\d+) points").unwrap();
    let reader = advent_of_code::file::read();
    let string_ex = String::from(reader);
    let match_str = &regex_string.captures(&*string_ex).unwrap();
    let nb_players = match_str[1].parse::<usize>().unwrap();
    let max_marble = match_str[2].parse::<usize>().unwrap();
    let mut score = vec![0; nb_players];
    let mut list = vec![0];
    for index in 1..max_marble {
        if index % 23 == 0 {
            list.rotate_right(7);
            score[index % nb_players] += index + list.pop().unwrap();
            list.rotate_left(1);
        } else {
            list.rotate_left(1);
            list.push(index);
        }
    }
    println!("{:#?}", score.iter().max().unwrap());
}

fn part_two() {
    let regex_string = Regex::new(r"^(\d+) players; last marble is worth (\d+) points").unwrap();
    let reader = advent_of_code::file::read();
    let string_ex = String::from(reader);
    let match_str = &regex_string.captures(&*string_ex).unwrap();
    let nb_players = match_str[1].parse::<usize>().unwrap();
    let max_marble = match_str[2].parse::<usize>().unwrap() * 100;
    let mut score = vec![0; nb_players];
    let mut list = VecDeque::with_capacity(max_marble);
    list.push_back(0);
    for index in 1..max_marble {
        if index % 23 == 0 {
            for _ in 0..7 {
                let popped = list.pop_back().unwrap();
                list.push_front(popped);
            }
            score[index % nb_players] += index + list.pop_back().unwrap();
            let popped = list.pop_front().unwrap();
            list.push_back(popped);
        } else {
            let popped = list.pop_front().unwrap();
            list.push_back(popped);
            list.push_back(index);
        }
    }
    println!("{:#?}", score.iter().max().unwrap());
}