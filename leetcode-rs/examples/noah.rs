#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
impl Solution {
    pub fn max_result(nums: Vec<i32>, k: i32) -> i32 {
        let n = nums.len();
        let mut dp = vec![0; n];
        dp[0] = nums[0];
        let mut q = std::collections::VecDeque::new();
        q.push_back(0usize);
        for i in 1..n {
            if i - q.front().unwrap() > k as usize {
                q.pop_front();
            }
            dp[i] = dp[*q.front().unwrap() as usize] + nums[i];
            while q.len() > 0 && dp[*q.back().unwrap()] <= dp[i] {
                q.pop_back();
            }
            q.push_back(i);
        }
        dp[n - 1]
    }
}

fn main() {}
