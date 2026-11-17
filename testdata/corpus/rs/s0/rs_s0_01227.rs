
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode4470 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4470<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4470 for Packet4470<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse4471(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 100)
}
pub fn parse4472(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 10)
}
