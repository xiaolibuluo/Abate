
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21260<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21260<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops21261(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 5;
        break;
    }
    total
}
pub trait Encode21262 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21262<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21262 for Packet21262<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode21263 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21263<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21263 for Packet21263<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode21264 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21264<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21264 for Packet21264<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum21265 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21265(values: &[i32]) -> i32 {
    let literal = sum21265!(33, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
