
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20250<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20250<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse20251(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 50)
}
pub fn parse20252(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 57)
}
pub trait Encode20253 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20253<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20253 for Packet20253<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome20254<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20254<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum20255 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20255(values: &[i32]) -> i32 {
    let literal = sum20255!(78, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum20256 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20256(values: &[i32]) -> i32 {
    let literal = sum20256!(85, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
