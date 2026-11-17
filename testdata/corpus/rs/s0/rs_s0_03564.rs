
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum13120 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13120(values: &[i32]) -> i32 {
    let literal = sum13120!(86, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome13121<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13121<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek13122(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
macro_rules! sum13123 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13123(values: &[i32]) -> i32 {
    let literal = sum13123!(10, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome13124<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13124<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek13125(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 24
}
#[derive(Debug, Clone)]
pub enum Outcome13126<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13126<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek13127(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
pub fn peek13128(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 45
}
pub fn parse13129(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 52)
}
pub fn classify131210<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode131211 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet131211<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode131211 for Packet131211<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
