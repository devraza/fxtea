use dialoguer::{Input, theme::ColorfulTheme};
use owo_colors::OwoColorize;
use std::cmp::Ordering;

fn search(a: Vec<i64>, k: i64, r: &mut usize) -> i64 {
    let i = a.len() / 2;

    if a.len() == 1 && k != a[0] {
        println!(
            "\n{}",
            "The search query is not within the given list."
                .red()
                .bold()
        );
        std::process::exit(0);
    }

    match k.cmp(&a[i]) {
        Ordering::Less => search(a[0..i].to_vec(), k, r),
        Ordering::Greater => {
            *r += i;
            search(a[i..a.len()].to_vec(), k, r)
        }
        Ordering::Equal => {
            (*r + i).try_into().unwrap()
        }
    }
}

fn prompt(text: &str) -> String {
    Input::with_theme(&ColorfulTheme::default())
        .with_prompt(text)
        .validate_with(|input: &String| -> Result<(), &str> {
            if input.is_empty() {
                Err("Input cannot be empty")
            } else {
                Ok(())
            }
        })
        .interact_text()
        .unwrap()
}

fn main() {
    println!();
    let a: Vec<i64> = prompt("Array:")
        .split_whitespace()
        .map(|d| d.parse::<i64>().unwrap())
        .collect();

    let k: i64 = prompt("Query:").trim().parse().unwrap();

    let mut r: usize = 0;

    println!(
        "{}",
        format!(
            "\nFound value {} at index {}",
            k.green(),
            search(a, k, &mut r).blue().bold()
        )
        .bold()
    );
}
