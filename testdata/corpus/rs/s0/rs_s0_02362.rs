
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek8900(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
#[derive(Debug, Clone)]
pub enum Outcome8901<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8901<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse8902(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 46)
}
pub fn peek8903(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 53
}
macro_rules! sum8904 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8904(values: &[i32]) -> i32 {
    let literal = sum8904!(60, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
