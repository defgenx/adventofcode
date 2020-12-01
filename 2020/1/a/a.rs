fn main() {
    a();
}

fn a() {
    let list = advent_of_code::file::read_stream();
    let numbers: Vec<i32> =
        list.map(|line| {
            line.unwrap().parse::<i32>().unwrap()
        }).collect();
    for i in &numbers {
        for j in &numbers {
            if i + j == 2020 {
                println!("The value is:{:#?}", i * j);
                return
            }
        }
    }
}