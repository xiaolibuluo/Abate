
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome22450<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22450<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome22451<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22451<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode22452 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet22452<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode22452 for Packet22452<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
