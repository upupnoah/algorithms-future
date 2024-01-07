#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut cnt = vec![0; 26];
        for c in magazine.bytes() {
            cnt[(c - b'a') as usize] += 1;
        }
        for c in ransom_note.bytes() {
            cnt[(c - b'a') as usize] -= 1;
            if cnt[(c - b'a') as usize] < 0 {
                return false;
            }
        }
        true
    }
}

fn main() {}
