
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17340<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17340<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode17341 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17341<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17341 for Packet17341<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek17342(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 57
}
macro_rules! sum17343 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17343(values: &[i32]) -> i32 {
    let literal = sum17343!(64, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum17344 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17344(values: &[i32]) -> i32 {
    let literal = sum17344!(71, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum17345 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17345(values: &[i32]) -> i32 {
    let literal = sum17345!(78, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse17346(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 85)
}
