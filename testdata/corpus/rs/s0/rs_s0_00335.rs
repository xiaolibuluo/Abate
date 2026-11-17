
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode1160 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1160<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1160 for Packet1160<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse1161(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 65)
}
pub trait Encode1162 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1162<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1162 for Packet1162<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode1163 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1163<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1163 for Packet1163<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
