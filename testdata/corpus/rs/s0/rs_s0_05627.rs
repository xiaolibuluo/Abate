
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse20370(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
#[derive(Debug, Clone)]
pub enum Outcome20371<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20371<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek20372(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
pub trait Encode20373 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20373<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20373 for Packet20373<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse20374(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
pub trait Encode20375 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20375<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20375 for Packet20375<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome20376<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20376<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek20377(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 54
}
pub trait Encode20378 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20378<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20378 for Packet20378<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20379 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20379(values: &[i32]) -> i32 {
    let literal = sum20379!(68, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
