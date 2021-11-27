use cargo_snippet::snippet;

fn main() {
    println!("Hello, world!");
}

#[snippet]
fn gcd(a: u64, b: u64) ->u64 {
    if b == 0 {
        a
    } else {
        gcd(b, a % b)
    }
}
