#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn repeat_limited_string(s: String, repeat_limit: i32) -> String {
        let mut cnt = [0; 26];
        for &v in s.as_bytes() {
            cnt[(v - b'a') as usize] += 1;
        }
        let mut ans = Vec::new();
        let mut i = 25;
        while i < 26 {
            let mut k = if i == 0 { 26 } else { i - 1 };
            loop {
                for _ in 0..repeat_limit {
                    if cnt[i] > 0 {
                        cnt[i] -= 1;
                        ans.push(b'a' + i as u8);
                    }
                }
                if cnt[i] == 0 {
                    break;
                }
                while k < 26 && cnt[k] == 0 {
                    if k == 0 {
                        k = 26;
                    } else {
                        k -= 1;
                    }
                }
                if k == 26 {
                    break;
                }
                cnt[k] -= 1;
                ans.push(b'a' + k as u8);
            }
            if i == 0 {
                break;
            }
            i -= 1;
        }
        String::from_utf8(ans).unwrap()
    }
}

fn main() {}
