#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn is_winner(player1: Vec<i32>, player2: Vec<i32>) -> i32 {
        let s1 = score(&player1);
        let s2 = score(&player2);
        match s1.cmp(&s2) {
            std::cmp::Ordering::Greater => 1,
            std::cmp::Ordering::Less => 2,
            std::cmp::Ordering::Equal => 0,
        }
    }
}

fn score(player: &Vec<i32>) -> i32 {
    let mut res = 0;
    let mut st = 0;
    player.into_iter().for_each(|x| {
        res += *x;
        if st > 0 {
            res += *x;
            st -= 1;
        }
        if *x == 10 {
            st = 2;
        }
    });
    res
}

fn main() {}
