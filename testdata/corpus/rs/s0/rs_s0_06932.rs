
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek24960(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 55
}
#[derive(Debug, Clone)]
pub enum Outcome24961<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24961<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome24962<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24962<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum24963 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24963(values: &[i32]) -> i32 {
    let literal = sum24963!(76, 68);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
