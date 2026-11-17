
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse21500(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 19)
}
pub fn peek21501(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
pub fn parse21502(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
pub trait Encode21503 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21503<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21503 for Packet21503<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
