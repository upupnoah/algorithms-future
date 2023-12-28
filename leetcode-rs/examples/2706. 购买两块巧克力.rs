#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn buy_choco(prices: Vec<i32>, money: i32) -> i32 {
        let (mut s1, mut s2) = (std::i32::MAX, std::i32::MAX);
        // prices.into_iter().for_each(|x| {
        //     if x < s1 {
        //         s2 = s1;
        //         s1 = x;
        //     } else if x < s2 {
        //         s2 = x;
        //     }
        // });
        for x in prices {
            if x < s1 {
                s2 = s1;
                s1 = x;
            } else if x < s2 {
                s2 = x;
            }
        }
        if s1 + s2 > money {
            money
        } else {
            money - s1 - s2
        }
    }
}

fn main() {}
