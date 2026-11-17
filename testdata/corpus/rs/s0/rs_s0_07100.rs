
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome25580<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25580<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode25581 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25581<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25581 for Packet25581<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome25582<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25582<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek25583(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 9
}
#[derive(Debug, Clone)]
pub enum Outcome25584<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25584<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum25585 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25585(values: &[i32]) -> i32 {
    let literal = sum25585!(23, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek25586(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 30
}
