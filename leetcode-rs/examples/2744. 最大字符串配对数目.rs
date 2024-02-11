#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
impl Solution {
    pub fn maximum_number_of_string_pairs(words: Vec<String>) -> i32 {
        let mut ans = 0;
        let mut st = vec![vec![false; 26]; 26];
        for word in words {
            let (x, y) = (
                (word.as_bytes()[0] - b'a') as usize,
                (word.as_bytes()[1] - b'a') as usize,
            );
            if st[y][x] {
                ans += 1;
            } else {
                st[x][y] = true;
            }
        }
        ans
    }
}

fn main() {}
