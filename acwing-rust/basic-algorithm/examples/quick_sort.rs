use std::io;

fn quick_sort(q: &mut [i32], l: i32, r: i32) {
    if l >= r {
        return;
    }
    let x = q[(l + r) as usize >> 1];
    let mut i = l - 1;
    let mut j = r + 1;
    while i < j {
        i += 1;
        while q[i as usize] < x {
            i += 1;
        }
        j -= 1;
        while q[j as usize] > x {
            j -= 1;
        }
        if i < j {
            q.swap(i as usize, j as usize);
        }
    }
    quick_sort(q, l, j);
    quick_sort(q, j + 1, r);
}

fn main() {
    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("Failed to read input");

    let n: i32 = input.trim().parse().expect("Invalid input");

    let mut q = String::new();
    io::stdin().read_line(&mut q).expect("Failed to read input");

    let q: Vec<i32> = q
        .split_whitespace()
        .map(|s| s.parse().expect("Invalid input"))
        .collect();

    let mut q = q.into_boxed_slice();
    quick_sort(&mut q, 0, n - 1);

    for i in 0..n {
        print!("{} ", q[i as usize]);
    }
}
