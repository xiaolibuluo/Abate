
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome13030<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13030<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum13031 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13031(values: &[i32]) -> i32 {
    let literal = sum13031!(73, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse13032(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 80)
}
pub fn peek13033(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
pub fn peek13034(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
pub fn classify13035<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode13036 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13036<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13036 for Packet13036<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek13037(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 18
}
pub fn parse13038(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 25)
}
pub fn peek13039(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
