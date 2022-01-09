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
    let max = x[0];
    let a = x[1];
    let b = x[2];

    let ans = (1..max + 1)
        .filter(|n| {
            println!("hello: {}", n);
            let sum = n
                .to_string()
                .chars()
                .map(|c| {
                    println!("{}", c as u8 - b'0');
                    (c as u8 - b'0') as u32
                })
                .sum::<u32>();
            a <= sum && sum <= b
        })
        .sum::<u32>();
    println!("{}", ans);
}
