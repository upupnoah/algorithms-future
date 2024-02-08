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
    // 使用二叉树的 LCA 解题, (一个二叉搜索树自然也是二叉树)
    pub fn lowest_common_ancestor1(
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

    // 利用二叉搜索树的性质(BST), 左子树一定< or <= 当前节点, 右子树一定>= or > 当前当前节点
    // 由于本题值都是唯一的, 因此不需要考虑是< or <= (都可以)
    pub fn lowest_common_ancestor(
        mut root: Option<Rc<RefCell<TreeNode>>>,
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        match (p, q) {
            (Some(p), Some(q)) => {
                while let Some(r) = root {
                    if r.borrow().val > p.borrow().val && r.borrow().val > q.borrow().val {
                        root = r.borrow_mut().left.take();
                    } else if r.borrow().val < p.borrow().val && r.borrow().val < q.borrow().val {
                        root = r.borrow_mut().right.take();
                    } else {
                        return Some(r);
                    }
                }
            }
            _ => None,
        }
    }
}
fn main() {}
