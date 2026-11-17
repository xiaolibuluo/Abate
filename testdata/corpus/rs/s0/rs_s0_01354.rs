
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode4970 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4970<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4970 for Packet4970<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome4971<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4971<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse4972(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 78)
}
