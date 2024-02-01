#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn stone_game_vi(alice_values: Vec<i32>, bob_values: Vec<i32>) -> i32 {
        let n = alice_values.len();
        let mut sum = vec![0; n];
        for i in 0..n {
            sum[i] = alice_values[i] + bob_values[i];
        }
        let mut idx = (0..n).collect::<Vec<_>>();
        idx.sort_by_key(|&i| -sum[i]);
        let (mut alice, mut bob) = (0, 0);
        for i in 0..n {
            if i & 1 == 0 {
                alice += alice_values[idx[i]];
            } else {
                bob += bob_values[idx[i]];
            }
        }
        alice.cmp(&bob) as i32
    }
}

fn main() {}
