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
    let a = x[0];
    let b = x[1];
    let ans = a * b;
    println!("{}", ans);
    if ans % 2 != 0 {
        println!("Odd");
    } else {
        println!("Even");
    }
}