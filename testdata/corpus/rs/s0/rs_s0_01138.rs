
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4140<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4140<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode4141 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4141<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4141 for Packet4141<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum4142 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4142(values: &[i32]) -> i32 {
    let literal = sum4142!(66, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse4143(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 73)
}
pub fn classify4144<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 80 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum4145 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4145(values: &[i32]) -> i32 {
    let literal = sum4145!(87, 34);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
