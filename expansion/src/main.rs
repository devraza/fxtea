use std::io;

const EXPONENTS: &str = "⁰¹²³⁴⁵⁶⁷⁸⁹";

fn choose(n: i64, r: i64) -> i64 {
    if r > n-r {
        return choose(n, n-r);
    }
    let mut result = 1;
    for i in 0..r {
        result = result * (n - i) / (i + 1);
    }
    result
}

fn superscript(n: i64) -> String {
    let n = n.to_string();

    let mut combined = String::new();
    for c in n.chars() {
        if c == '-' {
            combined.push_str("-");
        } else {
            let pos = c.to_digit(10).unwrap();
            combined.push(EXPONENTS.chars().nth(pos as usize).unwrap())
        }
    }

    combined
}

fn main() {
    println!("Enter a, b, and k, separated by a space:");
    let mut s = String::new();
    io::stdin().read_line(&mut s).unwrap();

    let nums: Vec<i64> = s.split_whitespace().map(|d| d.parse::<i64>().unwrap()).collect();

    let a = nums[0];
    let b = nums[1];
    let k = nums[2];

    let mut pow_a: u32 = k.try_into().unwrap();
    let mut pow_b: u32 = 0;

    let mut expansion: Vec<String> = vec![];

    for i in 0..k+1 {
        let mut end = String::new();
        if pow_b != k as u32 {
            end = format!("x{}", superscript(pow_a.into()));
            pow_a -= 1;
        }

        expansion.push(format!("{} {}", choose(k as i64, i as i64) * (a.pow(pow_a) * b.pow(pow_b)) as i64, end));

        pow_b += 1;
    }

    println!("{}", expansion.join(" + "));
}
