use std::{
    fs::File,
    io::{prelude::*, BufReader},
};

fn main() {
    problem_1();
}

fn count_increases(numbers: &Vec<i64>, window: usize) -> i64 {
    let mut count = 0;

    for (pos, _elt) in numbers.iter().enumerate() {
        let end = pos + 1;

        if pos < window || end > numbers.len() {
            continue;
        }

        let previous: &i64 = &numbers[pos - window..pos].iter().sum();
        let current: &i64 = &numbers[pos - window + 1..end].iter().sum();

        if current > previous {
            count += 1
        }
    }

    count
}

fn problem_1() {
    let filename = "input";

    let file = File::open(filename).expect("file wasn't found.");
    let reader = BufReader::new(file);

    let numbers: Vec<i64> = reader
        .lines()
        .map(|line| line.unwrap().parse::<i64>().unwrap())
        .collect();

    println!("problem 1a: {}", count_increases(&numbers, 1));
    println!("problem 1b: {}", count_increases(&numbers, 3));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_count_increases_window_size_1() {
        let numbers = vec![0, 1, 2];
        assert!(2 == count_increases(&numbers, 1));
    }

    #[test]
    fn test_count_increases_window_size_3() {
        let numbers = vec![0, 1, 2, 3, 4, 5];
        assert!(3 == count_increases(&numbers, 3));
    }
}
