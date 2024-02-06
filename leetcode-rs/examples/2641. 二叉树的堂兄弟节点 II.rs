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
impl Solution {
    pub fn replace_value_in_tree(
        root: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let mut queue = std::collections::VecDeque::new();
        if let Some(node) = root.clone() {
            node.borrow_mut().val = 0;
            queue.push_back(node.clone());
        }
        while !queue.is_empty() {
            let t = queue.clone();
            queue.clear();
            let mut next_level_sum = 0;
            for node in &t {
                if let Some(left) = node.borrow().left.clone() {
                    queue.push_back(left.clone());
                    next_level_sum += left.borrow().val;
                }
                if let Some(right) = node.borrow().right.clone() {
                    queue.push_back(right.clone());
                    next_level_sum += right.borrow().val;
                }
            }
            for node in t {
                let mut cur = 0;
                if let Some(left) = node.borrow().left.clone() {
                    cur += left.borrow().val;
                }
                if let Some(right) = node.borrow().right.clone() {
                    cur += right.borrow().val;
                }
                if let Some(left) = node.borrow().left.clone() {
                    left.borrow_mut().val = next_level_sum - cur;
                }
                if let Some(right) = node.borrow().right.clone() {
                    right.borrow_mut().val = next_level_sum - cur;
                }
            }
        }
        root
    }
}

fn main() {}
