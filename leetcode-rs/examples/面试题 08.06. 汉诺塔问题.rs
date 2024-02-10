#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
impl Solution {
    pub fn hanota(a: &mut Vec<i32>, b: &mut Vec<i32>, c: &mut Vec<i32>) {
        let n = a.len() as i32;
        Self::dfs(n, a, b, c);
    }
    fn dfs(n: i32, a: &mut Vec<i32>, b: &mut Vec<i32>, c: &mut Vec<i32>) {
        if n == 1 {
            // 考虑只有一个盘子, 那么直接将 a 中的盘子移动到 c 即可
            c.push(a.pop().unwrap());
            return;
        }
        Self::dfs(n - 1, a, c, b); // 先将 n - 1 个盘子从 a 移动到 b (借助 c)
        Self::dfs(1, a, b, c); // 将 a 中的最后一个盘子移动到 c
        Self::dfs(n - 1, b, a, c); // 最后将 n - 1 个盘子从 b 移动到 c (借助 a)
    }
}

fn main() {}
