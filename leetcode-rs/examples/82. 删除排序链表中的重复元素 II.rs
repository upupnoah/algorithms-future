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
// impl Solution {
//     pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
//         let mut res = Some(Box::new(ListNode::new(0)));
//         let mut p = res.as_mut().unwrap();
//         let mut pre = 101;
//         while let Some(mut node) = head {
//             head = node.next.take();
//             if (head.is_some() && head.as_ref().unwrap().val == node.val) || node.val == pre {
//                 pre = node.val;
//             } else {
//                 pre = node.val;
//                 p.next = Some(node);
//                 p = p.next.as_mut().unwrap();
//             }
//         }
//         res.as_mut().unwrap().next.take()
//     }
// }
impl Solution {
    pub fn delete_duplicates(mut head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode::new(0)); // 创建一个哑节点作为新链表的头部

        let mut tail = &mut dummy; // `tail` 用于跟踪新链表的最后一个节点

        let mut prev_val = None; // `prev_val` 用于记录前一个节点的值

        while let Some(mut node) = head {
            head = node.next.take(); // 取出当前节点，并将 `head` 移动到下一个节点
            let node_val = node.val; // 在潜在移动 `node` 之前，存储当前节点的值
            if head.as_ref().map_or(true, |next| next.val != node_val)
                && prev_val.map_or(true, |val| val != node_val)
            {
                tail.next = Some(node); // 如果当前节点是唯一的（值不等于下一个节点或前一个节点的值），则将其添加到新链表中

                tail = tail.next.as_mut().unwrap(); // 更新 `tail` 为新链表的最后一个节点
            }

            prev_val = Some(node_val); // 更新 `prev_val` 为当前节点的值
        }
        dummy.next
    }
}
fn main() {}
