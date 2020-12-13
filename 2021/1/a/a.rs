fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let mut count_incr: i32 = 0;
    let mut buff: i32 = 0;
    let numbers: Vec<i32> =
        list.map(|line| {
            line.unwrap().parse::<i32>().unwrap()
        }).collect();
    for i in &numbers {
        if buff == 0 {
            buff = *i;
            continue
        }
        if buff < *i {
            count_incr += 1;
        }
        buff = *i;
    }
    println!("The value is:{:#?}", count_incr);
}