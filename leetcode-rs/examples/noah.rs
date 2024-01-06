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
    pub fn insert_greatest_common_divisors(
        mut head: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut cur = &mut head;

        while cur.as_ref().unwrap().next.is_some() {
            let x = cur.as_mut().unwrap();
            let next = x.next.take();
            x.next = Some(Box::new(ListNode {
                val: gcd(x.val, next.as_ref().unwrap().val),
                next,
            }));
            cur = &mut cur.as_mut().unwrap().next.as_mut().unwrap().next;
        }
        head
    }
}
fn gcd(a: i32, b: i32) -> i32 {
    return if b == 0 { a } else { gcd(b, a % b) };
}

fn main() {}
