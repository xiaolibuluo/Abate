
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12250<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12250<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12251<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12251<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode12252 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12252<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12252 for Packet12252<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
