
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18570<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18570<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse18571(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 97)
}
pub trait Encode18572 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet18572<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode18572 for Packet18572<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
