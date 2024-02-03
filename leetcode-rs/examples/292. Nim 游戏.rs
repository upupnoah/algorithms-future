#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
impl Solution {
    pub fn can_win_nim(n: i32) -> bool {
        n % 4 != 0
    }
}

fn main() {}
