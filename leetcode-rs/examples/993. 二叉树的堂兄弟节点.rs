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
    pub fn is_cousins(root: Option<Rc<RefCell<TreeNode>>>, x: i32, y: i32) -> bool {
        type Tree = Option<Rc<RefCell<TreeNode>>>;
        fn dfs(root: Tree, parent: Tree, x: i32, depth: i32) -> Option<(i32, Tree)> {
            if let Some(node) = &root {
                let node = node.borrow();
                if node.val == x {
                    return Some((depth, parent));
                }
                // 如果为 None, 我还是得看看右边的情况, 否则说明找到了
                if let Some(l) = dfs(node.left.clone(), root.clone(), x, depth + 1) {
                    return Some(l);
                }
                return dfs(node.right.clone(), root.clone(), x, depth + 1);
            }
            None
        }
        let (dx, px) = dfs(root.clone(), None, x, 0).unwrap();
        let (dy, py) = dfs(root.clone(), None, y, 0).unwrap();
        dx == dy && px != py
    }
}
fn main() {}
