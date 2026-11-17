
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek21790(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
pub fn peek21791(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
#[derive(Debug, Clone)]
pub enum Outcome21792<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21792<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode21793 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21793<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21793 for Packet21793<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome21794<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21794<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome21795<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21795<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek21796(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
pub fn parse21797(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 57)
}
#[derive(Debug, Clone)]
pub enum Outcome21798<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21798<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops21799(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 71;
        break;
    }
    total
}
