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
use core::slice::sort;
use std::cell::RefCell;
use std::rc::Rc;
use std::slice::range;
type Tree = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    pub fn kth_largest_level_sum(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> i64 {
        let mut sums = vec![];
        let mut q = vec![];
        if let Some(r) = root {
            q.push(r);
            while !q.is_empty() {
                let mut t = vec![];
                let mut res = 0i64;
                for (i, v) in q.iter().enumerate() {
                    res += v.borrow().val as i64;
                    if let Some(l) = v.borrow().left.clone() {
                        t.push(l);
                    }
                    if let Some(r) = v.borrow().right.clone() {
                        t.push(r);
                    }
                }
                q = t;
                sums.push(res);
            }
        }
        if sums.len() < k as usize {
            -1
        }
        sums.sort_by(|a, b| b.cmp(a));
        sums[k as usize - 1]
    }
}

fn main() {}
