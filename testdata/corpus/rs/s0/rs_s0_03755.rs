
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum13770 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13770(values: &[i32]) -> i32 {
    let literal = sum13770!(58, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse13771(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 65)
}
#[derive(Debug, Clone)]
pub enum Outcome13772<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13772<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode13773 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13773<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13773 for Packet13773<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
