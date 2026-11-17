
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome600<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record600<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek601(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 16
}
pub trait Encode602 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet602<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode602 for Packet602<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode603 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet603<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode603 for Packet603<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
