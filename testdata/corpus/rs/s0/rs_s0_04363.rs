
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15980<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15980<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek15981(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 28
}
pub trait Encode15982 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15982<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15982 for Packet15982<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15983 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15983(values: &[i32]) -> i32 {
    let literal = sum15983!(42, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek15984(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 49
}
pub fn classify15985<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 56 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15986(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 63
}
pub fn parse15987(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 70)
}
#[derive(Debug, Clone)]
pub enum Outcome15988<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15988<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum15989 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15989(values: &[i32]) -> i32 {
    let literal = sum15989!(84, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
