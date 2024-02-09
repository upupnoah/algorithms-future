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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        type Tree = Option<Rc<RefCell<TreeNode>>>;
        let ans = Vec::new();
        fn dfs(root: Tree, ans: Vec<i32>) {
            if let Some(node) = &root {
                Self::dfs(node.borrow().left)
            }
        }
        dfs(root.clone(), &ans);
        ans
    }
}

fn main() {}
