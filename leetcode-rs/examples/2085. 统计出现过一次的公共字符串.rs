#![allow(dead_code)]
struct Solution;
use std::collections::HashMap;
impl Solution {
    pub fn count_words(words1: Vec<String>, words2: Vec<String>) -> i32 {
        let mut cnt1 = HashMap::new();
        let mut cnt2 = HashMap::new();
        // for word in words1 {
        //     *cnt1.entry(word).or_insert(0) += 1;
        // }
        // for word in words2 {
        //     *cnt2.entry(word).or_insert(0) += 1;
        // }
        words1
            .into_iter()
            .for_each(|x| *cnt1.entry(x).or_insert(0) += 1);
        words2
            .into_iter()
            .for_each(|x| *cnt2.entry(x).or_insert(0) += 1);

        let mut ans = 0;
        for (k, v) in cnt1 {
            if let Some(&u) = cnt2.get(&k) {
                if v == 1 && u == 1 {
                    ans += 1;
                }
            }
        }
        ans
    }
}

fn main() {}
