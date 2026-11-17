
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum10740 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10740(values: &[i32]) -> i32 {
    let literal = sum10740!(96, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10741 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10741(values: &[i32]) -> i32 {
    let literal = sum10741!(6, 35);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
