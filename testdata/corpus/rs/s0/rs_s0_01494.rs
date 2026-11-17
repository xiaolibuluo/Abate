
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum5470 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5470(values: &[i32]) -> i32 {
    let literal = sum5470!(35, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome5471<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5471<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5472(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 49)
}
macro_rules! sum5473 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5473(values: &[i32]) -> i32 {
    let literal = sum5473!(56, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
