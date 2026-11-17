
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek14420(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 30
}
pub fn parse14421(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
#[derive(Debug, Clone)]
pub enum Outcome14422<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14422<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14423(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 51
}
#[derive(Debug, Clone)]
pub enum Outcome14424<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14424<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum14425 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14425(values: &[i32]) -> i32 {
    let literal = sum14425!(65, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum14426 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14426(values: &[i32]) -> i32 {
    let literal = sum14426!(72, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome14427<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14427<T> {
    pub value: T,
    pub tag: u32,
}
