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

//fn main() {
//    let x: Vec<u32> = read_line();
//    let num = x[0];
//    let mut vec = Vec::new();
//    for _ in 0..num {
//        let n: Vec<u32> = read_line();
//        vec.push(n[0]);
//    }
//    vec.sort();
//    let mut cnt = 0;
//    for (i, _) in vec.iter().enumerate() {
//        if i < vec.len() - 1 {
//            if vec[i] < vec[i + 1] {
//                cnt += 1;
//            }
//        }
//    }
//    println!("{}", cnt + 1);
//}

// 別解
fn main() {
    let x: Vec<u32> = read_line();
    let num = x[0];
    let mut vec = Vec::new();
    for _ in 0..num {
        let n: Vec<u32> = read_line();
        vec.push(n[0]);
    }
    // ベクターソート
    vec.sort();
    // 重複削除
    vec.dedup();
    println!("{}", vec.len());
}
