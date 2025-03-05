use std::io;

fn fibonacci(i: u64, j: u64, n: &mut usize, a: &mut Vec<u64>) {
    while a.len() != *n {
        a.push(j);
        fibonacci(j, i+j, n, a);
    }
}

fn main() {
    let mut s = String::new();
    io::stdin().read_line(&mut s).unwrap();
    let mut n: usize = s.trim().parse().unwrap();

    let mut a: Vec<u64> = vec![1];

    fibonacci(1, 1, &mut n, &mut a);

    println!("{:?}", a)
}
