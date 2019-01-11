use std::collections::HashMap;

fn main() {
    let reader = advent_of_code::file::read();
    let mut hash_map = HashMap::new();
    hash_map.insert((0,0), true);
    let mut x = 0;
    let mut y = 0;
    for my_char in reader.chars() {
        match my_char {
            '^' => y += 1,
            'v' => y -= 1,
            '>' => x += 1,
            '<' => x -= 1,
            _   => println!("Tu ne connais pas le mot magique.")
        }
        if !hash_map.contains_key(&(x, y)) {
            hash_map.insert((x, y), true);
        }
    }
    println!("{:#?}", hash_map.len());



    let reader = advent_of_code::file::read();
    let mut hash_map = HashMap::new();
    hash_map.insert((0,0), true);
    let mut xy = (0, 0);
    let mut xy2 = (0, 0);
    let mut person = 1;
    for my_char in reader.chars() {
        let mut xy_tmp = (0,0);
        match person {
            1 => {
                xy_tmp = my_match(my_char, xy.0,xy.1);
                xy = xy_tmp;
            },
            -1 => {
                xy_tmp = my_match(my_char, xy2.0,xy2.1);
                xy2 = xy_tmp;
            },
            _ => println!("Aurevoir.")
        }
        if !hash_map.contains_key(&(xy_tmp.0, xy_tmp.1)) {
            hash_map.insert((xy_tmp.0, xy_tmp.1), true);
        }
        person *= -1;
    }
    println!("{:#?}", hash_map.len());
}

fn my_match(my_char: char, mut x: i32, mut y: i32) -> (i32, i32){
    match my_char {
        '^' => y += 1,
        'v' => y -= 1,
        '>' => x += 1,
        '<' => x -= 1,
        _   => println!("Tu ne connais pas le mot magique.")
    }
    (x,y)
}