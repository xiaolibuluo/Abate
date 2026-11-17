
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode24320 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24320<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24320 for Packet24320<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek24321(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 6
}
pub fn parse24322(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 13)
}
pub fn parse24323(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 20)
}
#[derive(Debug, Clone)]
pub enum Outcome24324<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24324<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode24325 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24325<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24325 for Packet24325<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome24326<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24326<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode24327 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet24327<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode24327 for Packet24327<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome24328<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24328<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum24329 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24329(values: &[i32]) -> i32 {
    let literal = sum24329!(62, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
