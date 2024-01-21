#![allow(dead_code)]

use std::collections::BinaryHeap;
struct Solution;
impl Solution {
    pub fn maximum_swap(num: i32) -> i32 {
        let mut nums = Vec::new();
        let mut num = num;
        let mut heap = BinaryHeap::new();

        // 分解数字并构建堆
        while num > 0 {
            let digit = num % 10;
            heap.push(digit);
            nums.push(digit);
            num /= 10;
        }

        // 遍历数字，找到第一个可以交换的位置
        let n = nums.len();
        for i in (0..n).rev() {
            if let Some(&top) = heap.peek() {
                if nums[i] == top {
                    heap.pop();
                } else {
                    // 找到最后一个等于堆顶元素的位置
                    let mut last = 0;
                    for j in (0..i).rev() {
                        if nums[j] == top {
                            last = j;
                            break;
                        }
                    }
                    // 交换
                    nums.swap(i, last);
                    break;
                }
            }
        }

        // 重建数字
        let mut ans = 0;
        for &digit in nums.iter().rev() {
            ans = ans * 10 + digit;
        }
        ans
    }
}

fn main() {}
