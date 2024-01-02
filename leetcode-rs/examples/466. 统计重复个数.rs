#![allow(dead_code)]
struct Solution;
use std::collections::HashMap;
impl Solution {
    pub fn get_max_repetitions(s1: String, n1: i32, s2: String, n2: i32) -> i32 {
        let mut cnt: Vec<i32> = Vec::new(); // store 匹配完每个 s1 之后是匹配了多少个s2 的字符
        let mut hash: HashMap<i32, i32> = HashMap::new(); // 哈希表快速判断余数有没有出现过

        // i 枚举匹配每个 s1，k 是在 s2 里转的指针
        let (mut i, mut k) = (0, 0);
        while i < n1 {
            for c in s1.chars() {
                // matched
                if c == s2.chars().nth(k as usize % s2.len()).unwrap() {
                    k += 1; // 往后走
                }
            }
            cnt.push(k); // 每个 s1 匹配完, 把 s2 的指针位置 k 存下来

            // 如果余数 k 对应的指针位置之前已经出现过了, 就说明出现了循环节
            if hash.contains_key(&(k % s2.len() as i32)) {
                // 计算: 循环节中一共有多少个 s1 -> 当前和上次是第几个 s1 之间的差
                let a = i - hash[&(k % s2.len() as i32)];
                // 计算: 循环节里能匹配多少个 s2 中的字符
                let b = k - cnt[hash[&(k % s2.len() as i32)] as usize];
                // 还剩多少个循环节 = 后面剩下的 s1 的个数/循环节里的 s1 的个数
                // 剩下的循环节个数 * 每个循环节能匹配的 s2 中的字符数 -> 剩下的循环节能匹配 s2 中的字符数
                let mut res = (n1 - i - 1) / a * b;
                // 把最后不完整的循环节模拟一遍, 找出 能 匹配多少字符
                // 先遍历不完整循环节的 s1 个数
                for _ in 0..(n1 - i - 1) % a {
                    // 再遍历每个 s1
                    for c in s1.chars() {
                        // 如果能和 s2 当前位置匹配
                        if c == s2.chars().nth(k as usize % s2.len()).unwrap() {
                            k += 1;
                        }
                    }
                }
                res += k;

                // 匹配的字符数 / s2 长度 / n2 -> 能匹配的大 S2 的数量
                return res / s2.len() as i32 / n2;
            }
            // 如果没有出现循环节, 标记 s2 的余数位置 k 对应的是循环节的第几个 s1
            // 这里也要取余, 不然没法判断出现循环节(判断是 s2 的同一位置)
            hash.insert(k % s2.len() as i32, i);
            i += 1;
        }
        // 没有出现循环节
        // 一个字符都匹配不了, 那么 cnt 是空的, 直接返回 0
        if cnt.is_empty() {
            return 0;
        }

        // 返回最后匹配了 s2 多少个字符, 除以 s2 的长度就是匹配了多少个 s2
        // 再除以 n2 就是匹配了多少个大 S2
        return *cnt.last().unwrap() / s2.len() as i32 / n2;
    }
}

fn main() {}
