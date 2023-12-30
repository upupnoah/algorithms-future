#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn day_of_year(date: String) -> i32 {
        let (mut day, mut month, mut year) = (0, 0, 0);
        let months = vec![
            31, // 1
            28, // 2
            31, // 3
            30, // 4
            31, // 5
            30, // 6
            31, // 7
            31, // 8
            30, // 9
            31, // 10
            30, // 11
            31, // 12
        ];
        for (i, s) in date.split('-').enumerate() {
            match i {
                0 => year = s.parse::<i32>().unwrap(),
                1 => month = s.parse::<i32>().unwrap(),
                2 => day = s.parse::<i32>().unwrap(),
                _ => {}
            }
        }
        if month > 2 && is_leap(year) {
            day += 1;
        }
        for i in 0..(month - 1) as usize {
            day += months[i];
        }
        day
    }
}

fn is_leap(year: i32) -> bool {
    year % 400 == 0 || year % 4 == 0 && year % 100 != 0
}

fn main() {}
