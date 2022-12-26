use std::fs;

fn read_file() -> String {
    let path = "./src/input";

    fs::read_to_string(path).expect(r#"Should have been able to read the file"#)
}

pub fn run() {
    let raw_data: String = read_file();
    let data: Vec<&str> = raw_data.split('\n').collect();

    let mut groups: Vec<Vec<i64>> = vec![];
    let mut group: Vec<i64> = vec![];
    let mut num: i64;

    for record in data {
        if !record.is_empty() {
            num = record.parse::<i64>().unwrap();
            group.push(num);
        } else {
            groups.push(group.clone());
            group.clear()
        }
    }
    let mut calories: Vec<i64> = vec![]; 
    for item in groups {
        calories.push(item.into_iter().sum());
    }
    calories.sort();
    let top_three = calories[calories.len()-3..].into_iter().sum::<i64>();
    println!("{:?}", top_three);
}
