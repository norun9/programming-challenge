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
    let mut vec: Vec<u32> = read_line();
    vec.sort_by(|x, y| x.cmp(y).reverse());
    let mut alice = 0;
    let mut bob = 0;
    for (i, &x) in vec.iter().enumerate() {
        if i % 2 == 0 {
            alice += x
        } else {
            bob += x
        }
    }
    println!("{}", alice - bob);
}
