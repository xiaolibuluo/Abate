
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode14250 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14250<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14250 for Packet14250<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse14251(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 10)
}
pub fn peek14252(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
#[derive(Debug, Clone)]
pub enum Outcome14253<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14253<T> {
    pub value: T,
    pub tag: u32,
}
