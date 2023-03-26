use std::{fs};

fn read_file() -> Result<String, String> {
    let path = "./src/input";

    match fs::read_to_string(path) {
        Ok(file) => Ok(file),
        Err(e) => Err(format!("Unable to open file: {}", e)),
    } 
}

enum Move {
    Rock,
    Paper,
    Scisors,
}

pub fn run() {
    let f = read_file();

    println!("{}", f);
}