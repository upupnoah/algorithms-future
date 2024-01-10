#![allow(dead_code)]

use std::collections::VecDeque;
struct Solution;
impl Solution {
    pub fn min_length(s: String) -> i32 {
        let mut stk: VecDeque<char> = VecDeque::new();
        for c in s.chars() {
            match stk.back() {
                Some(&'A') if c == 'B' => {
                    stk.pop_back();
                }
                Some(&'C') if c == 'D' => {
                    stk.pop_back();
                }
                _ => {
                    stk.push_back(c);
                }
            }
        }
        stk.len() as i32
    }
}

fn main() {
    let s = "ABCFCD";
    Solution::min_length(s.to_string());
}
