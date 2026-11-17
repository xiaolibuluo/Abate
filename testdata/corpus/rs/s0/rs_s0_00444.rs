
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome1540<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1540<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode1541 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1541<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1541 for Packet1541<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode1542 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1542<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1542 for Packet1542<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
