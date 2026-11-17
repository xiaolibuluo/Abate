
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome2990<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2990<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum2991 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2991(values: &[i32]) -> i32 {
    let literal = sum2991!(19, 39);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek2992(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
