
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12240<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12240<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum12241 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12241(values: &[i32]) -> i32 {
    let literal = sum12241!(16, 48);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome12242<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12242<T> {
    pub value: T,
    pub tag: u32,
}
