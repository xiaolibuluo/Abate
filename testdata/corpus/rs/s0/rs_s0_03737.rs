
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome13700<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13700<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13701<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13701<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum13702 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13702(values: &[i32]) -> i32 {
    let literal = sum13702!(78, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
