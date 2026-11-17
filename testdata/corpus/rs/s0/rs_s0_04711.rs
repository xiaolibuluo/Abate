
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum17030 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17030(values: &[i32]) -> i32 {
    let literal = sum17030!(28, 35);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
