
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode23960 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23960<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23960 for Packet23960<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome23961<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23961<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse23962(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 30)
}
pub fn parse23963(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
