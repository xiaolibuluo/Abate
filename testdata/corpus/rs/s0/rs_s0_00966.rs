
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3510<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3510<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode3511 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3511<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3511 for Packet3511<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3512<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3512<T> {
    pub value: T,
    pub tag: u32,
}
