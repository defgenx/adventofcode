fn main() {
    part_one();
}

fn part_one() {
    let reader = advent_of_code::file::read_stream();
    let mut counter = 0;
    for line in reader {
        let my_line = line.unwrap();
        if count_vowels(my_line.to_owned()) >= 3 && find_double(my_line.to_owned()) && contain_pattern(my_line.to_owned()) {
            counter += 1;
        }
    }

    println!("{}", counter);
}

fn count_vowels(line: String) -> i32 {
    let sum = line.chars().fold(0, |sum, x| {
        if x == 'a' || x == 'e' || x == 'i' || x == 'o' || x == 'u' {
            sum + 1
        } else {
            sum
        }
    });
    sum
}

fn find_double(line: String) -> bool {
    let mut current_char: char = ' ';
    let mut found = false;
    for chara in line.chars() {
        if current_char == chara {
            found = true;
            break;
        }
        current_char = chara;
    }
    found
}

fn contain_pattern(line: String) -> bool {
    !line.contains("ab") && !line.contains("cd") && !line.contains("pq") && !line.contains("xy")
}