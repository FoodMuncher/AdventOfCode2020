use std::fs::File;
use std::io::{prelude::*, BufReader};

const TREE: char = '#';

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    println!("Part One: {}", count_trees(3, 1))
}

fn part_two() {
    const MOVES: [[usize; 2]; 5] = [
        [1, 1],
        [3, 1],
        [5, 1],
        [7, 1],
        [1, 2],
    ];
    let mut tree_count: u32 = 1;

    for i in 0..5 {
        tree_count = tree_count * count_trees(MOVES[i][0], MOVES[i][1]);
    }

    println!("Part Two: {}", tree_count);
}

fn count_trees(right_moves: usize, down_moves: usize) -> u32 {
    let file = File::open("trees.txt").expect("");
    let reader = BufReader::new(file);
    
    let mut tree_count: u32 = 0;
    let mut col: usize = 0;
    let mut row: usize = 0;  

    for line in reader.lines() {
        row = row + 1;
        if row%down_moves == 1 {
            continue
        }

        if line.unwrap().chars().nth(col).unwrap() == TREE {
            tree_count = tree_count + 1
        }
        
        col = modulo(col+right_moves, 31)
    }

    return tree_count;
}

fn modulo(a: usize, b: usize) -> usize {
    return ((a % b) + b) % b
}