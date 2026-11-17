
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum25350 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25350(values: &[i32]) -> i32 {
    let literal = sum25350!(77, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum25351 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25351(values: &[i32]) -> i32 {
    let literal = sum25351!(84, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
