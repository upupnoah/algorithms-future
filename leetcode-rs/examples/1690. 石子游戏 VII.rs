#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn stone_game_vii(stones: Vec<i32>) -> i32 {
        let n = stones.len();
        let mut dp = vec![vec![0; n]; n];
        for i in (0..n-1).rev() {
            let mut sum = stones[i];
            for j in i+1..n {
                sum += stones[j];
                dp[i][j] = std::cmp::max(sum-stones[i] - dp[i+1][j], sum-stones[j] - dp[i][j-1]);
            }
        }
        dp[0][n-1]
    }
}

fn main() {}