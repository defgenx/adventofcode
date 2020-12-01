fn main() {
    b();
}

fn b() {
    let list = advent_of_code::file::read_stream();
    let numbers: Vec<i32> =
        list.map(|line| {
            line.unwrap().parse::<i32>().unwrap()
        }).collect();
    for i in &numbers {
        for j in &numbers {
            for k in &numbers {
                if i + j + k == 2020 {
                    println!("The value is: {:#?}", i * j * k);
                    return
                }
            }
        }
    }
}