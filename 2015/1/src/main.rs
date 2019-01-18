extern crate advent_of_code;

fn main() {
    let data = advent_of_code::file::read();
    (count(data.clone()));
    (count_stop(data.clone()));
}

fn count(content: String) {
    let mut sum = 0;
    for char_val in content.chars() {
        if char_val == '(' {
            sum += 1;
        } else {
            sum -= 1;
        }
    }
    println!("{:#?}", sum)
}

fn count_stop(content: String) {
    let mut sum = 0;
    let mut pos: usize = 0;
    for (i, char_val) in content.chars().enumerate() {
        if char_val == '(' {
            sum += 1;
        } else {
            sum -= 1;
        }
        if sum == -1 {
            pos = i;
            break;
        }
    }
    println!("{:#?}", pos)
}

