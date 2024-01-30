#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn distinct_difference_array(nums: Vec<i32>) -> Vec<i32> {
        let n = nums.len();
        let mut diff = vec![0; n];
        let mut st = std::collections::HashSet::new();
        for (i, &v) in nums.iter().enumerate().rev() {
            diff[i] = st.len() as i32;
            st.insert(v);
        }
        st.clear();
        for (i, &v) in nums.iter().enumerate() {
            st.insert(v);
            diff[i] = st.len() as i32 - diff[i];
        }
        diff
    }
}

fn main() {}
