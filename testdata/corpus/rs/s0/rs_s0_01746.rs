
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek6430(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
pub trait Encode6431 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6431<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6431 for Packet6431<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse6432(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 36)
}
pub trait Encode6433 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6433<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6433 for Packet6433<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
