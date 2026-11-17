
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse8880(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 6)
}
#[derive(Debug, Clone)]
pub enum Outcome8881<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8881<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome8882<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8882<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum8883 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8883(values: &[i32]) -> i32 {
    let literal = sum8883!(27, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
