use std::fs::File;
use std::io::{prelude::*, BufReader};

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let file = File::open("passwords.txt").expect("");
    let reader = BufReader::new(file);

    let mut correct_count: u32 = 0;

    for line in reader.lines() {
        let line = line.expect("");
        let mut line = line.split_whitespace();

        let mut min_max = line.next().expect("").split("-");

        let min: u32 = min_max.next().expect("").parse().expect("");
        let max: u32 = min_max.next().expect("").parse().expect("");

        let expected_letter = line.next().expect("").chars().next().expect("");

        let mut letter_count: u32 = 0;

        for letter in line.next().expect("").chars() {
            if letter == expected_letter {
                letter_count = letter_count + 1;
            }
        }

        if (letter_count >= min) && (letter_count <= max) {
            correct_count = correct_count + 1;
        }
    }

    println!("Part One: {}", correct_count)
}

fn part_two() {
    let file = File::open("passwords.txt").expect("");
    let reader = BufReader::new(file);

    let mut correct_count: u32 = 0;

    for line in reader.lines() {
        let line = line.expect("");
        let mut line = line.split_whitespace();

        let mut min_max = line.next().expect("").split("-");

        let position1: usize = min_max.next().expect("").parse().expect("");
        let position2: usize = min_max.next().expect("").parse().expect("");

        let letter = line.next().expect("").chars().next().expect("");

        let mut letters = line.next().expect("").chars();

        let letter1 = letters.nth(position1 - 1).expect("");
        let letter2 = letters.nth(position2 - position1 - 1).expect("");

        if (letter1 == letter) && (letter2 != letter) {
            correct_count = correct_count + 1;
        } 
        else if (letter2 == letter) && (letter1 != letter) {
            correct_count = correct_count + 1;
        }
    }

    println!("Part Two: {}", correct_count)
}