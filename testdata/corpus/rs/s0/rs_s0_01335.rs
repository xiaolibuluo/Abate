
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4910(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 83
}
pub fn parse4911(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 90)
}
pub trait Encode4912 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4912<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4912 for Packet4912<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
