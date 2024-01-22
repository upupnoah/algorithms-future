#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn alternating_subarray(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let (mut ans, mut idx) = (-1, 0);
        for i in 1..n {
            let len = (i - idx + 1) as i32;
            if nums[i] - nums[idx] == (len - 1) % 2 {
                ans = ans.max(len)
            } else {
                if nums[i] - nums[i - 1] == 1 {
                    idx = i - 1;
                    ans = ans.max(2);
                } else {
                    idx = i;
                }
            }
        }
        ans
    }
}

fn main() {}
