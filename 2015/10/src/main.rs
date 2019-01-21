fn main() {
    ex(40);
    ex(50);
}

fn ex(nb_times: i32) {
    let number_file: &str = &advent_of_code::file::read();
    let mut number_tmp: Vec<char> = number_file.to_owned().chars().collect();
    let mut number: String = String::new();
    for _ in 0..nb_times {
        let mut total_same_number = 1;
        let mut current_number = &number_tmp[0];
        for j in &number_tmp[1..] {
            if j == current_number {
                total_same_number += 1;
            } else {
                number.push_str(&*total_same_number.to_string());
                number.push_str(&current_number.to_string());
                total_same_number = 1;
                current_number = j;
            }
        }
        number.push_str(&*total_same_number.to_string());
        number.push_str(&current_number.to_string());
        number_tmp = number.chars().collect();
        number = String::new();
    }
    println!("{:#?}", number_tmp.len());
}

// Slow => Q&D
//fn ex(nb_times: i32) {
//    let max = advent_of_code::file::read();
//    let mut max_tmp: String = max.to_owned();
//    let mut number: String = String::new();
//    for _ in 0..nb_times {
//        let mut total_max: usize = 0;
//        while total_max != max_tmp.len() {
//            let new_max = nb_number((&max_tmp[total_max..]).to_string());
//            number.push_str(&*new_max.to_string());
//            let current_val: String = String::from(&max_tmp[total_max..total_max+1]);
//            number.push_str(&*current_val);
//            total_max += new_max;
//        }
//        max_tmp = number.to_owned();
//        number = String::new();
//    }
//    println!("{}", max_tmp.len());
//}
//
//fn nb_number(number: String) -> usize {
//    let mut count = 1;
//    let mut tmp_char = ' ';
//    let num_enum = number.chars().enumerate();
//    for (index, value) in num_enum {
//        if index == 0 {
//            tmp_char = value;
//            continue;
//        }
//        if value == tmp_char {
//            count += 1;
//        } else {
//            break;
//        }
//    }
//    return count;
//}