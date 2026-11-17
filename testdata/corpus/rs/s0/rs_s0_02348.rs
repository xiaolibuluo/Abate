
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome8810<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8810<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode8811 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet8811<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode8811 for Packet8811<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
