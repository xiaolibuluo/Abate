
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome2600<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2600<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse2601(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 94)
}
pub fn classify2602<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode2603 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2603<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2603 for Packet2603<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
