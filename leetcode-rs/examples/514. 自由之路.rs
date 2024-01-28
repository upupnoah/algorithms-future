#![allow(dead_code)]

struct Solution;
impl Solution {
    pub fn find_rotate_steps(ring: String, key: String) -> i32 {
        let (n, m) = (ring.len(), key.len());
        let (ring, key) = (
            ring.chars().collect::<Vec<_>>(),
            key.chars().collect::<Vec<_>>(),
        );
        // pos 用于记录 ring 中每个字符出现的所有位置
        let mut pos = vec![vec![]; 26];
        for (i, &c) in ring.iter().enumerate() {
            pos[(c as u8 - b'a') as usize].push(i);
        }
        // dp[i][j] 表示拼写 key 的前 i 个字符，ring 的指针在 j 位置时的最小操作次数
        let mut dp = vec![vec![std::i32::MAX; n]; m];

        // 初始化 dp[0][...]，处理 key 的第一个字符
        for &i in pos[(key[0] as u8 - b'a') as usize].iter() {
            dp[0][i] = i.min(n - i) as i32 + 1; // +1 是按下按钮的操作
        }

        for i in 1..m {
            // 当前字符在 ring 中的位置
            for &j in pos[(key[i] as u8 - b'a') as usize].iter() {
                // 上一个字符在 ring 中的位置
                for &k in pos[(key[i - 1] as u8 - b'a') as usize].iter() {
                    dp[i][j] = std::cmp::min(
                        dp[i][j],
                        dp[i - 1][k]
                            + std::cmp::min((j + n - k) % n, (k + n - j) % n) as i32 // 旋转 ring 的最小距离
                            + 1, // +1 是按下按钮的操作
                    );
                }
            }
        }
        // 返回拼写完整个 key 时，ring 的指针在任意位置的最小操作次数
        *dp[m - 1].iter().min().unwrap()
    }
}
fn main() {}
