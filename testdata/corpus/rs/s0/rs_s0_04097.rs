
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15050<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15050<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek15051(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
pub trait Encode15052 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15052<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15052 for Packet15052<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse15053(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 94)
}
pub fn parse15054(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 101)
}
pub fn loops15055(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 11;
        break;
    }
    total
}
macro_rules! sum15056 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15056(values: &[i32]) -> i32 {
    let literal = sum15056!(18, 57);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome15057<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15057<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek15058(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
pub fn parse15059(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 39)
}
macro_rules! sum150510 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold150510(values: &[i32]) -> i32 {
    let literal = sum150510!(46, 40);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
