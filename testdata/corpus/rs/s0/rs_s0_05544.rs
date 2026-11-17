
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse20010(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 22)
}
pub trait Encode20011 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20011<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20011 for Packet20011<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
