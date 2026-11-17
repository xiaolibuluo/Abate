
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode21180 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21180<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21180 for Packet21180<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome21181<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21181<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse21182(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
#[derive(Debug, Clone)]
pub enum Outcome21183<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21183<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome21184<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21184<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum21185 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21185(values: &[i32]) -> i32 {
    let literal = sum21185!(26, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
