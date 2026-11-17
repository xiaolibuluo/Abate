
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode8190 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet8190<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode8190 for Packet8190<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek8191(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 86
}
pub fn peek8192(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 93
}
pub fn peek8193(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
pub trait Encode8194 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet8194<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode8194 for Packet8194<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome8195<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8195<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum8196 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8196(values: &[i32]) -> i32 {
    let literal = sum8196!(24, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse8197(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 31)
}
pub fn peek8198(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
macro_rules! sum8199 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8199(values: &[i32]) -> i32 {
    let literal = sum8199!(45, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
