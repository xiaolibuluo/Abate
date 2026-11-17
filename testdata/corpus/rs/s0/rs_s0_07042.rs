
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek25380(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
pub fn peek25381(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
macro_rules! sum25382 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25382(values: &[i32]) -> i32 {
    let literal = sum25382!(33, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
