
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum23610 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23610(values: &[i32]) -> i32 {
    let literal = sum23610!(46, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek23611(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 53
}
pub fn parse23612(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 60)
}
pub fn peek23613(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
#[derive(Debug, Clone)]
pub enum Outcome23614<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23614<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse23615(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 81)
}
#[derive(Debug, Clone)]
pub enum Outcome23616<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23616<T> {
    pub value: T,
    pub tag: u32,
}
