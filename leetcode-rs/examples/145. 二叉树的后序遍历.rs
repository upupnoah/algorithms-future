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
    pub fn postorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut res = vec![];
        Self::dfs(root.clone(), &mut res);
        res
    }
    fn dfs(root: Tree, res: &mut Vec<i32>) {
        if let Some(node) = root {
            Self::dfs(node.borrow().left.clone(), res);
            Self::dfs(node.borrow().right.clone(), res);
            res.push(node.borrow().val);
        }
    }

    pub fn postorder_traversal_stack(root: Tree) -> Vec<i32> {
        let mut stk = vec![];
        let mut res = vec![];
        let mut node = root;
        let mut prev = None;
        while node.is_some() || !stk.is_empty() {
            while let Some(n) = node {
                node = n.borrow_mut().left.take();
                stk.push(n)
            }
            if let Some(n) = stk.pop() {
                if n.borrow().right.is_none() || n.borrow().right == prev {
                    res.push(n.borrow().val);
                    prev = Some(n);
                    node = None;
                } else {
                    node = n.borrow_mut().right.take();
                    stk.push(n);
                }
            }
        }
        res
    }
}
fn main() {}
