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
    let a1: Vec<u32> = read_line();
    let a = a1[0];
    let b1: Vec<u32> = read_line();
    let b = b1[0];
    let c1: Vec<u32> = read_line();
    let c = c1[0];
    let x1: Vec<u32> = read_line();
    let x = x1[0];
    let mut ans = 0;
    for i in 0..a + 1 {
        for j in 0..b + 1 {
            for k in 0..c + 1 {
                if i * 50 + j * 100 + k * 500 == x {
                    ans += 1;
                }
            }
        }
    }
    println!("{}", ans);
}
