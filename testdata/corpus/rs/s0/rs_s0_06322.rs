
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek22780(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 34
}
#[derive(Debug, Clone)]
pub enum Outcome22781<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22781<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek22782(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
pub fn peek22783(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 55
}
pub fn peek22784(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 62
}
macro_rules! sum22785 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22785(values: &[i32]) -> i32 {
    let literal = sum22785!(69, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome22786<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22786<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode22787 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet22787<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode22787 for Packet22787<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum22788 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22788(values: &[i32]) -> i32 {
    let literal = sum22788!(90, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify22789<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 97 => name,
        Some(_) => "other",
        None => "empty",
    }
}
