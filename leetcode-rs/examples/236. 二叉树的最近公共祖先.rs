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
    pub fn lowest_common_ancestor(
        root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(x) = &root {
            if Rc::ptr_eq(x, p.as_ref().unwrap()) || Rc::ptr_eq(x, q.as_ref().unwrap()) {
                return root;
            }
            let (left, right) = (
                Self::lowest_common_ancestor(x.borrow().left.clone(), p.clone(), q.clone()),
                Self::lowest_common_ancestor(x.borrow().right.clone(), p.clone(), q.clone()),
            );
            if left.is_some() && right.is_some() {
                return root;
            }
            return if left.is_some() { left } else { right };
        }
        None
    }
}
fn main() {}
