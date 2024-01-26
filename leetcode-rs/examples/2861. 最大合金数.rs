#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn max_number_of_alloys(
        n: i32,
        _: i32,
        budget: i32,
        composition: Vec<Vec<i32>>,
        stock: Vec<i32>,
        cost: Vec<i32>,
    ) -> i32 {
        let mx = *stock.iter().min().unwrap() as i32 + budget / n;
        let check = |num| -> bool {
            composition.iter().any(|comp| {
                let mut money = 0 as i64;
                for (j, &v) in comp.iter().enumerate() {
                    if v as i64 * num > stock[j] as i64 {
                        money += (v as i64 * num - stock[j] as i64) * cost[j] as i64;
                    }
                }
                money <= budget as i64
            })
        };

        let (mut l, mut r) = (0, mx);
        while l < r {
            let mid = (l + r + 1) >> 1;
            if check(mid as i64) {
                l = mid
            } else {
                r = mid - 1
            }
        }
        l
    }
}

fn main() {}
