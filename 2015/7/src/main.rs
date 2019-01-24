extern crate regex;

use std::collections::HashMap;
use std::str::FromStr;

use regex::Regex;

const NOT: &str = "NOT";
const LSHIFT: &str = "LSHIFT";
const RSHIFT: &str = "RSHIFT";
const OR: &str = "OR";
const AND: &str = "AND";
const SIMPLE: &str = "SIMPLE";

#[derive(Clone,Debug)]
enum Operation {
    Simple(String, String),
    Not(String, String),
    Rshift(String, String, String),
    Lshift(String, String, String),
    Or(String, String, String),
    And(String, String, String),
}

impl Operation {
    fn values(&self) -> Vec<String> {
        let x= match self {
            Operation::Simple(x, y)
            | Operation::Not(x, y) => [x.clone(), y.clone()].to_vec(),
            Operation::Rshift(x, y, z)
            | Operation::Lshift(x, y, z)
            | Operation::Or(x, y, z)
            | Operation::And(x, y, z) => [x.clone(), y.clone(), z.clone()].to_vec(),
        };
        return x;
    }

    fn name(&self) -> &str{
        return match *self {
            Operation::Simple(_,_) => SIMPLE,
            Operation::Not(_,_) => NOT,
            Operation::Rshift(_,_,_) => RSHIFT,
            Operation::Lshift(_,_,_) => LSHIFT,
            Operation::Or(_,_,_) => OR,
            Operation::And(_,_,_) => AND
        }
    }
}

impl FromStr for Operation {
    type Err = ();
    fn from_str(s: &str) -> Result<Operation, ()> {
        let regex_string_not = Regex::new(r"^NOT ([a-z0-9]+)$").unwrap();
        let regex_string_others = Regex::new(r"^(.+) (LSHIFT|RSHIFT|OR|AND) (.+)$").unwrap();
        let split_str: Vec<_> = s.split(" -> ").collect();
        let left_str = split_str[0];
        let right_str = split_str[1];
        let mut match1: String = String::from("");
        let mut match2: String = String::from("");
        let mut op: &str = SIMPLE;
        let op_match: String;
        if left_str.contains(NOT) {
            let match_str = regex_string_not.captures(&*left_str).unwrap();
            match1 = match_str[1].parse().unwrap();
            op = NOT;
        } else if left_str.contains(LSHIFT)
            || left_str.contains(RSHIFT)
            || left_str.contains(OR)
            || left_str.contains(AND)
        {
            let match_str = regex_string_others.captures(&*split_str[0]).unwrap();
            match1 = match_str[1].parse().unwrap();
            match2 = match_str[3].parse().unwrap();
            op_match = match_str[2].parse().unwrap();
            op = &*op_match;
        }

        match op {
            SIMPLE => Ok(Operation::Simple(left_str.to_string(), right_str.to_string())),
            NOT => Ok(Operation::Not(match1, right_str.to_string())),
            RSHIFT => Ok(Operation::Rshift(match1, match2, right_str.to_string())),
            LSHIFT => Ok(Operation::Lshift(match1, match2, right_str.to_string())),
            OR => Ok(Operation::Or(match1, match2, right_str.to_string())),
            AND => Ok(Operation::And(match1, match2, right_str.to_string())),
            _ => Err(()),
        }
    }
}
use std::borrow::Borrow;
fn main() {
    let endpoint = String::from("a");
    parse();
//    part_one(endpoint);
}

fn parse() {
    let reader = advent_of_code::file::read_stream();
    let mut instruction_map: HashMap<_, _, _> = HashMap::new();
    let mut values: Vec<String>;
    let mut current_operation: Operation;
    let mut key: &str;
    let mut name: &str;
    for line in reader {
        let my_line = line.unwrap();
        current_operation = my_line.parse().clone().unwrap();
        values = current_operation.values();
        instruction_map.insert(String::from(values.last().unwrap().as_str()), values);
//        println!("{:#?}{:#?}", key, name);
    }
    println!("{:#?}", instruction_map);
}

//fn part_one(end_point: String) {}

fn is_numeric(captured: String) -> bool {
    captured.parse::<i32>().is_ok()
}