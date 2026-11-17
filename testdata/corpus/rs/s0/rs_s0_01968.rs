
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome7370<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7370<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode7371 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7371<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7371 for Packet7371<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
