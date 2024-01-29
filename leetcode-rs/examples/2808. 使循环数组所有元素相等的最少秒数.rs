#![allow(dead_code)]

struct Solution;
impl Solution {
    pub fn minimum_seconds(nums: Vec<i32>) -> i32 {
        let mut pos = std::collections::HashMap::new();
        nums.iter().enumerate().for_each(|(i, &x)| {
            pos.entry(x).or_insert(vec![]).push(i);
        });
        let n = nums.len();
        let mut ans = n;
        for p in pos.values() {
            let mut p = p.clone();
            p.push(p[0] + n);
            let mut mx = 0;
            for i in 1..p.len() {
                mx = mx.max((p[i] - p[i - 1]) / 2);
            }
            ans = ans.min(mx);
        }
        ans as i32
    }
}
fn main() {}
