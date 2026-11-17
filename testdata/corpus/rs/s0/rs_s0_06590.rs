
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode23770 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23770<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23770 for Packet23770<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify23771<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 67 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse23772(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 74)
}
pub fn parse23773(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 81)
}
#[derive(Debug, Clone)]
pub enum Outcome23774<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23774<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23775(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 95
}
#[derive(Debug, Clone)]
pub enum Outcome23776<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23776<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum23777 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23777(values: &[i32]) -> i32 {
    let literal = sum23777!(12, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek23778(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
macro_rules! sum23779 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23779(values: &[i32]) -> i32 {
    let literal = sum23779!(26, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
