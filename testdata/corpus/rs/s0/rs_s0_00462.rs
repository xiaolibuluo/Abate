
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome1640<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1640<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse1641(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 10)
}
pub fn peek1642(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
pub trait Encode1643 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1643<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1643 for Packet1643<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum1644 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1644(values: &[i32]) -> i32 {
    let literal = sum1644!(31, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
