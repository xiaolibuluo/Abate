
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15470<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15470<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15471 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15471<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15471 for Packet15471<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15472<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15472<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15473 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15473<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15473 for Packet15473<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15474 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15474(values: &[i32]) -> i32 {
    let literal = sum15474!(65, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse15475(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 72)
}
pub trait Encode15476 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15476<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15476 for Packet15476<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse15477(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 86)
}
