
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek20870(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 73
}
pub fn peek20871(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 80
}
pub trait Encode20872 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20872<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20872 for Packet20872<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn loops20873(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 94;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome20874<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20874<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum20875 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20875(values: &[i32]) -> i32 {
    let literal = sum20875!(11, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse20876(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 18)
}
#[derive(Debug, Clone)]
pub enum Outcome20877<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20877<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum20878 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20878(values: &[i32]) -> i32 {
    let literal = sum20878!(32, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops20879(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 39;
        break;
    }
    total
}
