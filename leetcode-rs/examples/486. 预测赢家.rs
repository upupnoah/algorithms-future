#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn predict_the_winner(nums: Vec<i32>) -> bool {
        let n = nums.len();
        let mut dp = vec![vec![0; n]; n];
        for i in (0..n-1).rev() {
            dp[i][i] = nums[i];
            for j in i+1..n {
                dp[i][j] = std::cmp::max(nums[i] - dp[i+1][j], nums[j] - dp[i][j-1]);
            }
        }
        dp[0][n-1] >= 0
    }
}

fn main() {}