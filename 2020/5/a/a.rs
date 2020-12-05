fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let mut highest_seat_id = 0;
    for line in list {
        let unwrapped_line = line.unwrap();
        let elts: Vec<char> = unwrapped_line.chars().collect();
        let fb_elts = elts[..7].to_vec();
        let row_nb = find_row(fb_elts);
        println!("Row: {:#?}", row_nb);
        let lr_elts = elts[7..].to_vec();
        let line_nb = find_line(lr_elts);
        println!("Line: {:#?}", line_nb);
        let new_highest_seat_id = row_nb * 8 + line_nb;
        println!("Seat ID: {:#?}", new_highest_seat_id);
        if highest_seat_id < new_highest_seat_id {
            highest_seat_id = new_highest_seat_id
        }
    }
    println!("Highest Seat ID: {:#?}", highest_seat_id);
}

fn find_row(parts: Vec<char>) -> i32 {
    let mut lower = 0;
    let mut upper = 127;
    for part in parts {
        println!("{:#?}", part);
        if part.to_string() == "F" {
            upper = compute_upper_right(upper, lower)
        } else {
            lower = compute_lower_left(lower, upper)
        }
        println!("{:#?}", lower);
        println!("{:#?}", upper);
    }
    return upper
}

fn find_line(parts: Vec<char>) -> i32 {
    let mut left = 0;
    let mut right = 7;
    for part in parts {
        if part.to_string() == "L" {
            right = compute_upper_right(right, left)
        } else {
            left = compute_lower_left(left, right)
        }
    }
    return right
}

fn compute_upper_right(a: i32, b: i32) -> i32 {
    return a - ((a - b) + 1) / 2
}

fn compute_lower_left(a: i32, b: i32) -> i32 {
    return a + ((b - a) + 1) / 2
}