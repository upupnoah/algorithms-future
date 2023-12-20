#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn is_acronym(words: Vec<String>, s: String) -> bool {
        words.len() == s.len()
            && words
                .into_iter()
                .zip(s.into_bytes())
                .all(|(w, c)| w.as_bytes()[0] == c)
    }
}

fn main() {}
