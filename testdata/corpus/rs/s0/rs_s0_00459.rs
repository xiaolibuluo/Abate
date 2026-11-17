
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome1620<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1620<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse1621(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 81)
}
pub fn peek1622(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 88
}
pub trait Encode1623 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1623<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1623 for Packet1623<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify1624<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 5 => name,
        Some(_) => "other",
        None => "empty",
    }
}
