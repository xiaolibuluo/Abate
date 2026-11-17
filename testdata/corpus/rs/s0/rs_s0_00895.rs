
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek3290(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 14
}
pub fn parse3291(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 21)
}
pub trait Encode3292 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3292<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3292 for Packet3292<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3293<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3293<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3294(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 42
}
pub fn parse3295(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 49)
}
macro_rules! sum3296 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3296(values: &[i32]) -> i32 {
    let literal = sum3296!(56, 60);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum3297 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3297(values: &[i32]) -> i32 {
    let literal = sum3297!(63, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome3298<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3298<T> {
    pub value: T,
    pub tag: u32,
}
