
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16190<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16190<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse16191(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 10)
}
pub trait Encode16192 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16192<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16192 for Packet16192<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse16193(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 24)
}
