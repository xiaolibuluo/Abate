
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek7760(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 5
}
pub fn loops7761(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 12;
        break;
    }
    total
}
pub fn peek7762(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
#[derive(Debug, Clone)]
pub enum Outcome7763<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7763<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse7764(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
macro_rules! sum7765 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7765(values: &[i32]) -> i32 {
    let literal = sum7765!(40, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome7766<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7766<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek7767(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
#[derive(Debug, Clone)]
pub enum Outcome7768<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7768<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum7769 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7769(values: &[i32]) -> i32 {
    let literal = sum7769!(68, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum77610 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold77610(values: &[i32]) -> i32 {
    let literal = sum77610!(75, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
