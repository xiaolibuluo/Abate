
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse15120(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 67)
}
pub trait Encode15121 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15121<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15121 for Packet15121<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
