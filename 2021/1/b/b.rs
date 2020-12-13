fn main() {
    b();
}

fn b() {
    let list = advent_of_code::file::read_stream();
    let mut count_incr: i32 = 0;
    let numbers: Vec<i32> =
        list.map(|line| {
            line.unwrap().parse::<i32>().unwrap()
        }).collect();
    for (i, _) in numbers.iter().enumerate() {
        let j = i + 1;
        if j+3 > (numbers.len() as i32).try_into().unwrap() {
            break
        }
        let a: Vec<i32> = numbers[i as usize..(i+3) as usize].to_vec();
        let b: Vec<i32> = numbers[j as usize..(j+3) as usize].to_vec();
        if (a.len() as i32) < 3 || (b.len() as i32) < 3 {
            break
        }
        if a.iter().sum::<i32>() < b.iter().sum::<i32>() {
            count_incr += 1;
        }
    }
    println!("The value is:{:#?}", count_incr);
}