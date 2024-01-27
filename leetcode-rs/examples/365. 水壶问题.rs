#![allow(dead_code)]
struct Solution;
impl Solution {
    // 贝祖定理
    pub fn can_measure_water1(x: i32, y: i32, t: i32) -> bool {
        if x + y < t {
            return false;
        }
        t % Self::gcd(x, y) == 0
    }

    fn gcd(a: i32, b: i32) -> i32 {
        if b == 0 {
            return a;
        }
        Self::gcd(b, a % b)
    }

    // DFS
    pub fn can_measure_water(x: i32, y: i32, t: i32) -> bool {
        let mut vis = std::collections::HashSet::new();
        Self::dfs(0, 0, x, y, t, &mut vis)
    }

    fn dfs(
        rx: i32,
        ry: i32,
        x: i32,
        y: i32,
        t: i32,
        vis: &mut std::collections::HashSet<(i32, i32)>,
    ) -> bool {
        if rx + ry == t {
            return true;
        }
        if vis.contains(&(rx, ry)) {
            return false;
        }
        vis.insert((rx, ry));
        // 倒满任意一个容器
        if Self::dfs(x, ry, x, y, t, vis) || Self::dfs(rx, y, x, y, t, vis) {
            return true;
        }
        // 清空任意一个容器
        if Self::dfs(0, ry, x, y, t, vis) || Self::dfs(rx, 0, x, y, t, vis) {
            return true;
        }
        // 从 x 倒入 y
        if Self::dfs(0.max(rx - (y - ry)), y.min(ry + rx), x, y, t, vis) {
            return true;
        }
        // 从 y 倒入 x
        if Self::dfs(x.min(rx + ry), 0.max(ry - (x - rx)), x, y, t, vis) {
            return true;
        }
        false
    }
}

fn main() {}
