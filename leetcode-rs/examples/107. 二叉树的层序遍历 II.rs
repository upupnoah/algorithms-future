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
use std::intrinsics::mir::Len;
use std::rc::Rc;
type Tree = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    pub fn level_order_bottom(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut q = vec![];
        let mut ans = vec![];
        if let Some(r) = root {
            q.push(r);
            while !q.is_empty() {
                let mut t = vec![];
                let mut res = vec![];
                for i in 0..q.len() {
                    res.push(q[i].borrow().val);
                    if let Some(left) = q[i].borrow_mut().left.take() {
                        t.push(left);
                    }
                    if let Some(right) = q[i].borrow_mut().right.take() {
                        t.push(right)
                    }
                }
                ans.push(res);
                q = t;
            }
        }
        ans.reverse();
        ans
    }
}

fn main() {}
