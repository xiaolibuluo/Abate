
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21990<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21990<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome21991<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21991<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum21992 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21992(values: &[i32]) -> i32 {
    let literal = sum21992!(88, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
