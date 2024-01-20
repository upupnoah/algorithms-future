#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn split_array(nums: Vec<i32>, m: i32) -> i32 {
        let n = nums.len();
        let mut f = vec![vec![std::i32::MAX; m as usize + 1]; n + 1];
        let mut sub = vec![0; n + 1];

        // 前缀和
        for i in 0..n {
            sub[i + 1] = sub[i] + nums[i];
        }
        f[0][0] = 0;
        for i in 1..=n {
            for j in 1..=i.min(m as usize) {
                // 枚举 k，其中前 k 个数被分割为 j−1 段，而第 k+1 到第 i 个数为第 j 段
                for k in 0..i {
                    f[i][j] = f[i][j].min(f[k][j - 1].max(sub[i] - sub[k]));
                }
            }
        }
        f[n][m as usize]
    }
}
fn main() {}
