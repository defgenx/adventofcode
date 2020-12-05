fn main() {
    b();
}

fn b() {
    let list = advent_of_code::file::read_stream();
    let mut seat_map: Vec<Vec<usize>> = vec![vec![0; 8]; 128];
    for line in list {
        let unwrapped_line = line.unwrap();
        let elts: Vec<char> = unwrapped_line.chars().collect();
        let fb_elts = elts[..7].to_vec();
        let row_nb = find_row(fb_elts) as usize;
        let lr_elts = elts[7..].to_vec();
        let line_nb = find_line(lr_elts) as usize;
        let seat_id: usize = (row_nb * 8 + line_nb) as usize;
        seat_map[row_nb][line_nb] = seat_id
    }
    let mut my_seat_id: usize = 0;
    for i in 0..128 {
        let row = i;
        for j in 0..8 {
            let line = j;
            print!("{:#?} | ", seat_map[row][line]);
            // We must add the first test only because we wanted to print the grid
            if line > 0 && (line+1) < 128 &&
                seat_map[row][line] == 0 &&
                seat_map[row][line-1] != 0 &&
                seat_map[row][line+1] != 0 {
                my_seat_id = row * 8 + line
            }
        }
        println!();
    }
    println!("My Seat ID: {:#?}", my_seat_id);
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