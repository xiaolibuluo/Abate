
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify14890<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome14891<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14891<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode14892 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14892<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14892 for Packet14892<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse14893(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 80)
}
macro_rules! sum14894 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14894(values: &[i32]) -> i32 {
    let literal = sum14894!(87, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek14895(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
macro_rules! sum14896 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14896(values: &[i32]) -> i32 {
    let literal = sum14896!(101, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
