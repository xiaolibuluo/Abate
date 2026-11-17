
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14260<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14260<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14261(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
pub trait Encode14262 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14262<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14262 for Packet14262<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome14263<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14263<T> {
    pub value: T,
    pub tag: u32,
}
