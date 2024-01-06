#![allow(dead_code)]

struct Solution;
impl Solution {
    // 单调栈
    // 从右往左遍历, 维护一个单调递减的栈
    // 如果当前元素比栈顶元素小, 那么当前元素就能看到栈顶元素
    pub fn can_see_persons_count(heights: Vec<i32>) -> Vec<i32> {
        let n = heights.len();
        let (mut ans, mut stk) = (vec![0; n], vec![std::i32::MAX]);
        for (i, v) in heights.into_iter().enumerate().rev() {
            while let Some(&top) = stk.last() {
                if top < v {
                    ans[i] += 1;
                    stk.pop();
                } else {
                    break;
                }
            }
            if stk.len() > 1 {
                ans[i] += 1;
            }
            stk.push(v);
        }
        ans
    }
}
fn main() {}
