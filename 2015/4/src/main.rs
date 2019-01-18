fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let reader = advent_of_code::file::read();
    let mut string_hash = String::from("");
    let string_ex = String::from(reader);
    let mut counter = 0;
    while !string_hash.starts_with("00000") {
        let concat = vec![string_ex.clone(), counter.to_string()].join("");
        string_hash = format!("{:x}", md5::compute(concat.clone()));
//        println!("{:#?}", md5::compute(format!("b\"{}\"", concat)));
        counter += 1;
    }
    println!("String {} has hash {} so response is {:#?}", string_ex, string_hash, (counter - 1));
}

fn part_two() {
    let reader = advent_of_code::file::read();
    let mut string_hash = String::from("");
    let string_ex = String::from(reader);
    let mut counter = 0;
    while !string_hash.starts_with("000000") {
        let concat = vec![string_ex.clone(), counter.to_string()].join("");
        string_hash = format!("{:x}", md5::compute(concat.clone()));
//        println!("{:#?}", md5::compute(format!("b\"{}\"", concat)));
        counter += 1;
    }
    println!("String {} has hash {} so response is {:#?}", string_ex, string_hash, (counter - 1));
}