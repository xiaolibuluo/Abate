
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome23870<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23870<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode23871 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23871<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23871 for Packet23871<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
