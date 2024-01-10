#![allow(dead_code)]
struct Solution;
impl Solution {
    // 方法一: 考虑 abc 的个数
    // 答案由 t 个 ‘abc’ 组成, 因此需要插入的字符个数最少是 3*t - len(s)
    // 如果s[i-1] >= s[i], 则需要生成一个新的 abc
    pub fn add_minimum(word: String) -> i32 {
        let mut t = 1;
        for i in 1..word.len() {
            if word.as_bytes()[i - 1] >= word.as_bytes()[i] {
                t += 1;
            }
        }
        t * 3 - word.len() as i32
    }

    // 方法二: 考虑相邻两个字母
    // ab  b-a-1 = 0 个 (不需要插入)
    // ac  c-a-1 = 1 个 (需要插入一个 b)
    // ca  a-c-1 (+3) = 0, 但是我们需要的 0(ca 之间不需要插入), 因此 + 3
    // 综上: (右-左+2) % 3
    // 坑点: 开头 和 结尾是没有相邻两个字符的, 因此需要特殊判断
    // s[0]: s[0]-'a'
    // s[n-1]: 'c'-s[n-1]
    // s[0] + s[n-1] = s[0]-s[n-1]+2
    pub fn add_minimum2(word: String) -> i32 {
        let mut ans = word.as_bytes()[0] - word.as_bytes()[word.len() - 1] + 2;
        for i in 1..word.len() {
            ans += (word.as_bytes()[i] - word.as_bytes()[i - 1] + 2) % 3;
        }
        ans as i32
    }
}

fn main() {}
