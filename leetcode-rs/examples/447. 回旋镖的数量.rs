#![allow(dead_code)]

use std::collections::HashMap;
struct Solution;
impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        let mut ans = 0;
        let mut cnt = HashMap::new();
        for p1 in &points {
            cnt.clear();
            for p2 in &points {
                let d = dist(p1, p2);
                let v = cnt.entry(d).or_insert(0);
                ans += *v * 2;
                *v += 1;
            }
        }
        ans
    }
}

fn dist(a: &Vec<i32>, b: &Vec<i32>) -> i32 {
    let dx = a[0] - b[0];
    let dy = a[1] - b[1];
    dx * dx + dy * dy
}

fn main() {}
