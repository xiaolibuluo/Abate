
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14160<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14160<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode14161 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet14161<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode14161 for Packet14161<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
