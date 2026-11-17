
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome10320<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10320<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome10321<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10321<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome10322<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10322<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode10323 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10323<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10323 for Packet10323<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
