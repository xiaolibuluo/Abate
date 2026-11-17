
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome26240<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26240<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops26241(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 77;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome26242<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26242<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode26243 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26243<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26243 for Packet26243<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek26244(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
pub fn peek26245(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
pub fn parse26246(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 15)
}
pub fn peek26247(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
pub fn parse26248(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 29)
}
pub fn loops26249(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 36;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome262410<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record262410<T> {
    pub value: T,
    pub tag: u32,
}
