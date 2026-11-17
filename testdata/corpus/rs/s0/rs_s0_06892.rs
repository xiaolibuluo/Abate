
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum24840 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24840(values: &[i32]) -> i32 {
    let literal = sum24840!(93, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome24841<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24841<T> {
    pub value: T,
    pub tag: u32,
}
