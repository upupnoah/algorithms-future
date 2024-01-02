#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn min_operations_max_profit(
        customers: Vec<i32>,
        boarding_cost: i32,
        running_cost: i32,
    ) -> i32 {
        let (mut ans, mut mx, mut profit, mut waiting, mut i) = (-1, 0, 0, 0, 0);
        while waiting > 0 || i < customers.len() {
            waiting += if i < customers.len() { customers[i] } else { 0 };
            let up = std::cmp::min(4, waiting);
            waiting -= up;
            i += 1;
            profit += up * boarding_cost - running_cost;

            if profit > mx {
                mx = t;
                ans = i as i32;
            }
        }
        ans
    }
}

fn main() {}
