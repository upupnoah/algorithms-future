#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn minimum_removal(mut beans: Vec<i32>) -> i64 {
        beans.sort();
        let (mut sum, mut mx) = (0 as i64, 0 as i64);
        beans.iter().enumerate().for_each(|(i, &v)| {
            sum += v as i64;
            mx = mx.max((beans.len() - i) as i64 * v as i64);
        });
        sum - mx
    }
}

fn main() {}
