
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome26010<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26010<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode26011 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26011<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26011 for Packet26011<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode26012 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26012<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26012 for Packet26012<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse26013(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 83)
}
pub fn peek26014(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 90
}
pub trait Encode26015 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26015<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26015 for Packet26015<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
