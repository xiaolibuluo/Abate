
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4300(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 66
}
#[derive(Debug, Clone)]
pub enum Outcome4301<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4301<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum4302 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4302(values: &[i32]) -> i32 {
    let literal = sum4302!(80, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode4303 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4303<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4303 for Packet4303<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek4304(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
#[derive(Debug, Clone)]
pub enum Outcome4305<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4305<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek4306(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 11
}
#[derive(Debug, Clone)]
pub enum Outcome4307<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4307<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek4308(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 25
}
#[derive(Debug, Clone)]
pub enum Outcome4309<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4309<T> {
    pub value: T,
    pub tag: u32,
}
