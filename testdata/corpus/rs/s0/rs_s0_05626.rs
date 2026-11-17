
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20360<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20360<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome20361<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20361<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify20362<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 6 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode20363 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20363<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20363 for Packet20363<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek20364(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 20
}
pub trait Encode20365 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20365<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20365 for Packet20365<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20366 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20366(values: &[i32]) -> i32 {
    let literal = sum20366!(34, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum20367 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20367(values: &[i32]) -> i32 {
    let literal = sum20367!(41, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
