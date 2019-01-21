fn main() {
    part_one();
}

fn part_one() {
    let max = advent_of_code::file::read();
    let mut max_tmp: String = max.to_owned();
    let mut number: String = String::new();
    for _ in 0..40 {
        let mut total_max: usize = 0;
        while total_max != max_tmp.len() {
            let new_max = nb_number((&max_tmp[total_max..]).to_string());
            number.push_str(&*new_max.to_string());
            let current_val: String = String::from(&max_tmp[total_max..total_max+1]);
            number.push_str(&*current_val);
            total_max += new_max;
        }
        max_tmp = number.to_owned();
        number = String::new();
    }
    println!("{}", max_tmp.len());
}

fn nb_number(number: String) -> usize {
    let mut count = 1;
    let mut tmp_char = ' ';
    let num_enum = number.chars().enumerate();
    for (index, value) in num_enum {
        if index == 0 {
            tmp_char = value;
            continue;
        }
        if value == tmp_char {
            count += 1;
        } else {
            break;
        }
    }
    return count;
}