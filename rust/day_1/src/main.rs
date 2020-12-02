use std::fs::File;
use std::io::{prelude::*, BufReader};
use std::collections::HashSet;

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let file = File::open("expense_report.txt").expect("");
    let reader = BufReader::new(file);

    let mut set: HashSet<u32> = HashSet::new();

    for line in reader.lines() {
        let line: u32 = line.expect("").parse().expect("");
        
        if set.contains(&(2020 - line)) {
            println!("Part One: {}", (2020 - line) * line);
            break;
        }

        set.insert(line);
    }
}

fn part_two() {
    let file = File::open("expense_report.txt").expect("");
    let reader = BufReader::new(file);

    let mut set: HashSet<i32> = HashSet::new();

    for line1 in reader.lines() {
        let line1: i32 = line1.expect("").parse().expect("");

        for line2 in set.iter() {
            if set.contains(&(2020 - line1 - line2)) {
                println!("Part Two: {}", (2020 - line1 - line2) * line1 * line2);
                break;
            }
        }

        set.insert(line1);
    }
}