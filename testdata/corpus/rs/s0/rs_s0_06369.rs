
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum22940 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22940(values: &[i32]) -> i32 {
    let literal = sum22940!(48, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek22941(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 55
}
#[derive(Debug, Clone)]
pub enum Outcome22942<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22942<T> {
    pub value: T,
    pub tag: u32,
}
