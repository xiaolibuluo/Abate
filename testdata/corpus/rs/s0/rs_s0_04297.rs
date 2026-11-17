
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse15740(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 97)
}
#[derive(Debug, Clone)]
pub enum Outcome15741<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15741<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15742 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15742<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15742 for Packet15742<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek15743(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 21
}
pub fn peek15744(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 28
}
