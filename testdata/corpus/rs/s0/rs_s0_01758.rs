
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome6480<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record6480<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek6481(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
macro_rules! sum6482 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6482(values: &[i32]) -> i32 {
    let literal = sum6482!(101, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome6483<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record6483<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse6484(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 18)
}
