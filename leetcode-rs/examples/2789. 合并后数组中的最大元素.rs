#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
impl Solution {
    pub fn max_array_value(nums: Vec<i32>) -> i64 {
        let mut ans = *nums.last().unwrap() as i64;
        for i in (0..nums.len() - 1).rev() {
            if ans >= nums[i - 1] as i64 {
                ans += nums[i - 1] as i64;
            } else {
                ans = nums[i - 1] as i64;
            }
        }
        ans
    }
}

fn main() {}
