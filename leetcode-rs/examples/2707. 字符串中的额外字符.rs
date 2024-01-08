#![allow(dead_code)]

use std::collections::HashSet;
struct Solution;
// dfs(i): 表示 s 的前 i 个字符的子问题
// 跳过 s 的最后一个字符, 有 dfs(i) = dfs(i-1)+1
// 考虑 [枚举选哪个], 如果从s[j] 到 s[i] 的子串在dictionary 中, 有 dfs(i) = min dfs(j-1)  (j = 0,1,...i)
// 递归边界: dfs(-1) = 0
// ans: dfs(n-1)

impl Solution {
    pub fn min_extra_char(s: String, dictionary: Vec<String>) -> i32 {
        let has = dictionary.into_iter().collect::<HashSet<_>>();
        let n = s.len();
        let mut memo = vec![-1; n];

        fn dfs(i: i32, memo: &mut [i32], has: &HashSet<String>, s: &str) -> i32 {
            if i < 0 {
                return 0;
            }

            if memo[i as usize] != -1 {
                return memo[i as usize];
            }

            // Don't choose the i-th character
            let mut res = dfs(i - 1, memo, has, s) + 1;

            // Enumerate which character to choose
            for j in 0..=i {
                if has.contains(&s[j as usize..=(i as usize)]) {
                    res = res.min(dfs(j - 1, memo, has, s));
                }
            }

            memo[i as usize] = res;
            res
        }

        dfs(n as i32 - 1, &mut memo, &has, &s)
    }
}
fn main() {}
