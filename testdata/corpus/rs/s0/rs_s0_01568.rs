
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum5760 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5760(values: &[i32]) -> i32 {
    let literal = sum5760!(24, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode5761 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5761<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5761 for Packet5761<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek5762(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
pub fn parse5763(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 45)
}
#[derive(Debug, Clone)]
pub enum Outcome5764<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5764<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5765(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 59)
}
#[derive(Debug, Clone)]
pub enum Outcome5766<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5766<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5767(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 73)
}
pub fn peek5768(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
pub fn peek5769(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
macro_rules! sum57610 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold57610(values: &[i32]) -> i32 {
    let literal = sum57610!(94, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
