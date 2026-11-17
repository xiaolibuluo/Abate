
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum10590 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10590(values: &[i32]) -> i32 {
    let literal = sum10590!(95, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse10591(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
macro_rules! sum10592 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10592(values: &[i32]) -> i32 {
    let literal = sum10592!(12, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome10593<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10593<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse10594(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 26)
}
pub trait Encode10595 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10595<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10595 for Packet10595<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek10596(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 40
}
pub trait Encode10597 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10597<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10597 for Packet10597<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
