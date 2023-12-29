#![allow(dead_code)]
struct Solution;
impl Solution {
    pub fn day_of_the_week(day: i32, month: i32, year: i32) -> String {
        // region:    --- 模拟
        // let months = [
        //     31, // 1
        //     28, // 2
        //     31, // 3
        //     30, // 4
        //     31, // 5
        //     30, // 6
        //     31, // 7
        //     31, // 8
        //     30, // 9
        //     31, // 10
        //     30, // 11
        //     31, // 12
        // ];
        // let mut days = 0;
        // // 输入年份之前的年份的天数贡献
        // days += (year - 1971) * 365 + (year - 1969) / 4;
        //  // 输入年份中，输入月份之前的月份的天数贡献
        // for i in 0..month - 1 {
        //     days += months[i as usize];
        // }
        // if month >= 3 && (year % 400 == 0 || year % 4 == 0 && year % 100 != 0) {
        //     days += 1;
        // }
        // // 输入月份中的天数贡献
        // days += day;
        // return match days % 7 {
        //     0 => "Thursday".to_string(),
        //     1 => "Friday".to_string(),
        //     2 => "Saturday".to_string(),
        //     3 => "Sunday".to_string(),
        //     4 => "Monday".to_string(),
        //     5 => "Tuesday".to_string(),
        //     6 => "Wednesday".to_string(),
        //     _ => "".to_string(),
        // };
        // endregion: --- 模拟

        // region:    --- 基姆拉尔森公式
        // let days = [
        //     "Sunday",
        //     "Monday",
        //     "Tuesday",
        //     "Wednesday",
        //     "Thursday",
        //     "Friday",
        //     "Saturday",
        // ];

        // let (d, mut m, mut y) = (day, month, year);
        // if m < 3 {
        //     m += 12;
        //     y -= 1;
        // }
        // let week = (d + 2 * m + 3 * (m + 1) / 5 + y + y / 4 - y / 100 + y / 400 + 1) % 7;
        // endregion: --- 基姆拉尔森公式

        // region:    --- 蔡勒公式
        let days = [
            "Sunday",
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
        ];

        let (mut y, mut m, d) = (year, month, day);
        if m < 3 {
            m += 12;
            y -= 1;
        }
        let (c, y) = (y / 100, y % 100);
        let mut week = (y + y / 4 + c / 4 - 2 * c + 26 * (m + 1) / 10 + d - 1);
        week = (week % 7 + 7) % 7;
        // endregion: --- 蔡勒公式

        return days[week as usize].to_string();
    }
}
fn main() {}
