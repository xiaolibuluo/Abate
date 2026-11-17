
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum26320 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26320(values: &[i32]) -> i32 {
    let literal = sum26320!(77, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome26321<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26321<T> {
    pub value: T,
    pub tag: u32,
}
