
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome2450<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2450<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode2451 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2451<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2451 for Packet2451<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
