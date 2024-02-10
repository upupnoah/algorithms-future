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
    pub fn preorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = vec![];
        Self::dfs(root.clone(), &mut res);
        res
    }
    fn dfs(root: Tree, res: &mut Vec<i32>) {
        if let Some(node) = root {
            res.push(node.borrow().val);
            Self::dfs(node.borrow().left.clone(), res);
            Self::dfs(node.borrow().right.clone(), res);
        }
    }

    pub fn preorder_traversal_stack(root: Tree) -> Vec<i32> {
        let mut res = Vec::new();
        let mut stk = vec![];
        let mut node = root;
        while node.is_some() || stk.len() > 0 {}
    }
}

fn main() {}
