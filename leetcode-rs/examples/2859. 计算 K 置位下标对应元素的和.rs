#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn sum_indices_with_k_set_bits(nums: Vec<i32>, k: i32) -> i32 {
        let mut ans = 0;
        nums.into_iter().enumerate().for_each(|(i, x)| {
            let mut i = i as i32;
            let mut cnt = 0;
            while i > 0 {
                i -= lowbit(i);
                cnt += 1;
            }
            if cnt == k {
                ans += x;
            }
        });
        ans
    }
}
fn lowbit(x: i32) -> i32 {
    x & -x
}

fn main() {}
