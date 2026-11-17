
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4070<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4070<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek4071(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub fn parse4072(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 72)
}
pub trait Encode4073 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4073<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4073 for Packet4073<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode4074 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4074<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4074 for Packet4074<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
