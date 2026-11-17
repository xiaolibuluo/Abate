
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome5740<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5740<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek5741(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 5
}
pub fn loops5742(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 12;
        break;
    }
    total
}
pub fn peek5743(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
#[derive(Debug, Clone)]
pub enum Outcome5744<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5744<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum5745 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5745(values: &[i32]) -> i32 {
    let literal = sum5745!(33, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse5746(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 40)
}
pub trait Encode5747 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5747<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5747 for Packet5747<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse5748(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 54)
}
