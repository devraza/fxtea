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
            combined.push('-');
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

    let mut expansion: Vec<String> = vec![];

    for i in 0..k+1 {
        let mut end = String::new();

        if i != k {
            if pow_a != 1 {
                end = format!("x{}", superscript(pow_a.into()));
            } else {
                end = String::from("x");
            }
        }

        expansion.push(format!("{} {}", choose(k, i) * (a.pow(pow_a) * b.pow(i as u32)), end));

        if i != k {
            pow_a -= 1;
        }

    }

    println!("{}", expansion.join(" + "));
}
