
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode15680 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15680<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15680 for Packet15680<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse15681(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 26)
}
pub trait Encode15682 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15682<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15682 for Packet15682<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
