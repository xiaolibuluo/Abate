
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4920<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4920<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode4921 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4921<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4921 for Packet4921<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
