#![allow(dead_code)]
#![feature(trace_macros)]
struct Solution;
// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}
use std::cell::RefCell;
use std::rc::Rc;
type Tree = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut ans = vec![];

        // 也可以直接使用 vec
        let mut q = std::collections::VecDeque::new();

        if let Some(r) = root {
            q.push_back(r);
            while !q.is_empty() {
                let mut tmp = std::collections::VecDeque::new();
                let mut res = vec![];
                while !q.is_empty() {
                    let node = q.pop_front().unwrap(); // first take out.
                    let mut node = node.borrow_mut();
                    res.push(node.val);
                    if let Some(left) = node.left.take() {
                        tmp.push_back(left);
                    }
                    if let Some(right) = node.right.take() {
                        tmp.push_back(right);
                    }
                }
                ans.push(res);
                q = tmp;
            }
        }
        ans
    }
}

fn main() {}
