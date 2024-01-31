#![allow(dead_code)]

struct Solution;

use std::cmp::Reverse;
use std::collections::BinaryHeap;
impl Solution {
    pub fn nums_game(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut res = vec![0; n];
        let mut lower = BinaryHeap::new(); // 最大堆
        let mut upper = BinaryHeap::new(); // 最小堆，使用 Reverse 包装元素
        let mut lower_sum = 0i64;
        let mut upper_sum = 0i64;
        let mod_val = 1_000_000_007i64;

        for i in 0..n {
            let x = nums[i] - i as i32;
            if lower.is_empty() || *lower.peek().unwrap() >= x {
                lower.push(x);
                lower_sum += x as i64;
                if lower.len() > upper.len() + 1 {
                    let top = lower.pop().unwrap();
                    lower_sum -= top as i64;
                    upper.push(Reverse(top));
                    upper_sum += top as i64;
                }
            } else {
                upper.push(Reverse(x));
                upper_sum += x as i64;
                if lower.len() < upper.len() {
                    let top = upper.pop().unwrap().0;
                    upper_sum -= top as i64;
                    lower.push(top);
                    lower_sum += top as i64;
                }
            }
            if (i + 1) % 2 == 0 {
                res[i] = ((upper_sum - lower_sum) % mod_val) as i32;
            } else {
                res[i] = ((upper_sum - lower_sum + *lower.peek().unwrap() as i64) % mod_val) as i32;
            }
        }
        res
    }
}

fn main() {}
