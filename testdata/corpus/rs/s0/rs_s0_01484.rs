
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum5430 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5430(values: &[i32]) -> i32 {
    let literal = sum5430!(80, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum5431 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5431(values: &[i32]) -> i32 {
    let literal = sum5431!(87, 39);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
