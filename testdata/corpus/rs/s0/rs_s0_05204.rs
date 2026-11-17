
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18780<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18780<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops18781(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 79;
        break;
    }
    total
}
pub fn loops18782(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 86;
        break;
    }
    total
}
pub fn peek18783(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 93
}
pub fn parse18784(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 100)
}
#[derive(Debug, Clone)]
pub enum Outcome18785<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18785<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum18786 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18786(values: &[i32]) -> i32 {
    let literal = sum18786!(17, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
