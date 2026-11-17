
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome23040<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23040<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome23041<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23041<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23042(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 95
}
pub trait Encode23043 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23043<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23043 for Packet23043<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek23044(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 12
}
