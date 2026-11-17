
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome8130<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8130<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome8131<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8131<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek8132(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
pub fn loops8133(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 22;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome8134<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8134<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse8135(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 36)
}
macro_rules! sum8136 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8136(values: &[i32]) -> i32 {
    let literal = sum8136!(43, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome8137<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8137<T> {
    pub value: T,
    pub tag: u32,
}
