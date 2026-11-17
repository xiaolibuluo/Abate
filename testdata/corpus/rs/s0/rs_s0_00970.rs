
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3540<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3540<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome3541<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3541<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum3542 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3542(values: &[i32]) -> i32 {
    let literal = sum3542!(62, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome3543<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3543<T> {
    pub value: T,
    pub tag: u32,
}
