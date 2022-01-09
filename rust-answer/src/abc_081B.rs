use std::fmt::Debug;
use std::str::FromStr;

fn read_line<T>() -> Vec<T>
where
    T: FromStr,
    <T as FromStr>::Err: Debug,
{
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim()
        .split_whitespace()
        .map(|c| T::from_str(c).unwrap())
        .collect()
}

fn main() {
    let x: Vec<u32> = read_line();
    let _num = x[0];
    let mut nums: Vec<u32> = read_line();

    let mut count: u32 = 0;
    loop {
        // 全て偶数かどうか
        if !nums.iter().any(|&n| n % 2 != 0) {
            nums = nums.iter().map(|i| i / 2).collect::<Vec<u32>>();
            // println!("{:?}", nums);
            count += 1;
        } else {
            break;
        }
    }
    println!("{}", count);
}
