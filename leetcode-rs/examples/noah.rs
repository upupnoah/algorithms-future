#![allow(dead_code)]
struct Solution;
impl Solution {
    // 输入视角: 选/不选
    pub fn maximum_rows(matrix: Vec<Vec<i32>>, num_select: i32) -> i32 {
        let mut selected = vec![false; matrix[0].len()];

        fn check(matrix: &Vec<Vec<i32>>, selected: &Vec<bool>) -> i32 {
            let mut cnt = 0;
            for i in 0..matrix.len() {
                for j in 0..matrix[0].len() {
                    if !selected[j] && matrix[i][j] == 1 {
                        cnt -= 1;
                        break;
                    }
                }
                cnt += 1;
            }
            cnt
        }

        fn dfs(
            matrix: &Vec<Vec<i32>>,
            selected: &mut Vec<bool>,
            i: usize,
            cnt: i32,
            num_select: i32,
            ans: &mut i32,
        ) {
            if i >= matrix[0].len() {
                if cnt == num_select {
                    *ans = (*ans).max(check(matrix, selected));
                }
                return;
            }

            // 选第 i 列
            selected[i] = true;
            dfs(matrix, selected, i + 1, cnt + 1, num_select, ans);
            selected[i] = false;

            // 不选第 i 列
            dfs(matrix, selected, i + 1, cnt, num_select, ans);
        }

        let mut ans = 0;
        dfs(&matrix, &mut selected, 0, 0, num_select, &mut ans);
        ans
    }

    // 答案视角: 枚举选哪个
    pub fn maximum_rows2(matrix: Vec<Vec<i32>>, num_select: i32) -> i32 {
        fn check(matrix: &Vec<Vec<i32>>, selected: &Vec<bool>) -> i32 {
            let mut cnt = 0;
            for i in 0..matrix.len() {
                for j in 0..matrix[0].len() {
                    if !selected[j] && matrix[i][j] == 1 {
                        cnt -= 1;
                        break;
                    }
                }
                cnt += 1;
            }
            cnt
        }
        fn dfs(
            matrix: &Vec<Vec<i32>>,
            selected: &mut Vec<bool>,
            i: usize,
            cnt: i32,
            num_select: i32,
            ans: &mut i32,
        ) {
            if cnt == num_select {
                *ans = (*ans).max(check(matrix, selected));
            }
            for j in i..matrix[0].len() {
                selected[j] = true;
                dfs(matrix, selected, j + 1, cnt + 1, num_select, ans);
                selected[j] = false;
            }
        }
        let mut selected = vec![false; matrix[0].len()];
        let mut ans = 0;
        dfs(&matrix, &mut selected, 0, 0, num_select, &mut ans);
        ans
    }
}

fn main() {}
