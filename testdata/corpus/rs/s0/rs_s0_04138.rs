
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15210<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15210<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome15211<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15211<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15212 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15212<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15212 for Packet15212<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse15213(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 11)
}
