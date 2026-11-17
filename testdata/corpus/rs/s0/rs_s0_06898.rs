
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum24860 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24860(values: &[i32]) -> i32 {
    let literal = sum24860!(22, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum24861 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24861(values: &[i32]) -> i32 {
    let literal = sum24861!(29, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
