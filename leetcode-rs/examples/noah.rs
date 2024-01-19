#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn minimum_time(nums1: Vec<i32>, nums2: Vec<i32>, x: i32) -> i32 {
        let mut pairs = nums1.iter().zip(nums2.iter()).collect::<Vec<_>>();
        pairs.sort_unstable_by(|&a, &b| a.1.cmp(&b.1));

        let n = pairs.len();
        let mut f = vec![0; n + 1];
        for (i, &(a, b)) in pairs.iter().enumerate() {
            for j in (1..=i + 1).rev() {
                f[j] = f[j].max(f[j - 1] + a + b * j as i32);
            }
        }

        let s1 = nums1.iter().sum::<i32>();
        let s2 = nums2.iter().sum::<i32>();
        for (t, &v) in f.iter().enumerate() {
            if s1 + s2 * t as i32 - v <= x {
                return t as i32;
            }
        }
        -1
    }
}

fn main() {}
