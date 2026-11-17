
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome7900<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7900<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome7901<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7901<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode7902 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet7902<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode7902 for Packet7902<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
