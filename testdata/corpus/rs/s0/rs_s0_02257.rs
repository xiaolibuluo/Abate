
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse8400(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 61)
}
#[derive(Debug, Clone)]
pub enum Outcome8401<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8401<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode8402 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet8402<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode8402 for Packet8402<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome8403<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8403<T> {
    pub value: T,
    pub tag: u32,
}
