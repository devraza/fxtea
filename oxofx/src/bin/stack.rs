use std::io::{self, Write};

fn tail(stack: [i64; 5]) -> usize {
    let mut index = 0;
    for i in 0..=4 {
        let subtract = -1*i as isize;
        let wanted_index = (4+subtract) as usize;

        if subtract <= 0 && stack[wanted_index] != 0 {
            index = wanted_index;
            break;
        }
    }
    index
}

fn push(stack: &mut [i64; 5], number: &mut String) {
    let mut index = 0;
    let mut used = 0;

    for (idx, i) in stack.clone().iter().enumerate() {
        if *i == 0 {
            index = idx;
            break;
        } else {
            used += 1;
        }
    }
    if used < 5 {
        stack[index] = number.parse().unwrap();
    } else {
        eprintln!("\x1b[31m\x1b[1m\nThe stack is full!\x1b[0m");
    }
}

fn main() {
    let mut stack: [i64; 5] = [0; 5];

    println!("\x1b[32m\x1b[1m\n  Push, peek, pop or quit\x1b[0m");
    println!("  When pushing, append the value to push to the command: `push x`\n");
    let mut arguments: Vec<&str> = vec!["0"];
    let mut action = String::new();
    let mut argument: String;

    while arguments[0] != "quit" {
        print!("\x1b[36m$\x1b[0m ");
        io::stdout().flush().unwrap();

        io::stdin().read_line(&mut action).expect("Failed to read input");

        argument = action.clone().replace("\n", "");
        arguments = argument.split(' ').collect::<Vec<&str>>();

        action = String::new();

        match arguments[0] {
            "push" => {
                push(&mut stack, &mut arguments[1].to_string());
                println!("\n\x1b[1m{:?}\x1b[0m\n", stack);
            }
            "peek" => {
                println!("\n\x1b[1m{}\x1b[0m\n", stack[tail(stack)]);
            }
            "pop" => {
                let used: Vec<i64> = stack.iter().filter(|&&x| x != 0).cloned().collect();
                if used.len() > 0 {
                    stack[tail(stack)] = 0;
                    eprintln!("\n\x1b[1m{:?}\x1b[0m\n", stack);
                } else {
                    eprintln!("\x1b[31m\x1b[1m\nThe stack is empty!\x1b[0m\n");
                }
            }
            "quit" => {
                std::process::exit(0);
            }
            _ => {
                eprintln!("\x1b[31m\x1b[1m\nThis action is not supported!\x1b[0m\n");
            }
        }
    }
}
