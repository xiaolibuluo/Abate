
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode1200 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1200<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1200 for Packet1200<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode1201 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1201<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1201 for Packet1201<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode1202 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1202<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1202 for Packet1202<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse1203(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 34)
}
pub trait Encode1204 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1204<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1204 for Packet1204<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
