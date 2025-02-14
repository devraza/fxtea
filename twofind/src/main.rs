use owo_colors::OwoColorize;
use dialoguer::{Input, theme::ColorfulTheme};

fn search(a: Vec<i64>, k: i64, r: &mut usize) -> i64 {
    let i = a.len()/2;

    if a.len() == 1 {
        println!("\n{}", "The search query is not within the array. You have a massive skill issue.".red().bold());
        std::process::exit(0);
    }

    if k < a[i] {
        search(a[0..i].to_vec(), k, r)
    } else if k > a[i] {
        *r += i;
        search(a[i..a.len()].to_vec(), k, r)
    } else {
        return (*r+a.len()/2).try_into().unwrap();
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
    let a: Vec<i64> = prompt("Array:").split_whitespace().map(|d| d.parse::<i64>().unwrap()).collect();

    let k: i64 = prompt("Query:").trim().parse().unwrap();

    let mut r: usize= 0;

    println!("{}", format!("\nFound value {} at index {}", k.green(), search(a, k, &mut r).blue().bold()).bold());
}
