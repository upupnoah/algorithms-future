#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn maximum_sum_of_heights(max_heights: Vec<i32>) -> i64 {
        let n = max_heights.len();

        let mut stk = vec![n as i32]; // 哨兵
        let mut suf = vec![0; n + 1]; // 后缀最大值
        let mut sum = 0 as i64;

        // 计算后缀最大值
        max_heights.iter().enumerate().rev().for_each(|(i, &x)| {
            while stk.len() > 1 && x <= max_heights[*stk.last().unwrap() as usize] {
                let j = stk.pop().unwrap();
                sum -= max_heights[j as usize] as i64 * (*stk.last().unwrap() - j) as i64;
            }
            sum += x as i64 * (stk.last().unwrap() - i as i32) as i64;
            stk.push(i as i32);
            suf[i] = sum;
        });

        let mut ans = sum;

        sum = 0;
        stk = vec![-1];
        // 计算前缀最大值 & 计算答案
        max_heights.iter().enumerate().for_each(|(i, &x)| {
            while stk.len() > 1 && x <= max_heights[*stk.last().unwrap() as usize] {
                let j = stk.pop().unwrap();
                sum -= max_heights[j as usize] as i64 * (j - *stk.last().unwrap()) as i64;
            }
            sum += x as i64 * (i as i32 - *stk.last().unwrap()) as i64;
            stk.push(i as i32);
            ans = ans.max(sum as i64 + suf[i + 1] as i64);
        });

        ans
    }
}

fn main() {}
