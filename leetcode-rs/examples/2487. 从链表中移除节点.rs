#![allow(dead_code)]

struct Solution;
// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}
impl Solution {
    pub fn remove_nodes(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        if let Some(mut node) = head {
            let next = node.next.take();
            node.next = Solution::remove_nodes(next);
            if let Some(next_node) = node.next.as_ref() {
                if node.val < next_node.val {
                    return node.next;
                }
            }
            Some(node)
        } else {
            None
        }
    }
}

fn main() {}
