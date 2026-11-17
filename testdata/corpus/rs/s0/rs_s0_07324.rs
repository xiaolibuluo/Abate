
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome26420<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26420<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse26421(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 20)
}
#[derive(Debug, Clone)]
pub enum Outcome26422<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26422<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek26423(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 34
}
pub trait Encode26424 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet26424<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode26424 for Packet26424<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
