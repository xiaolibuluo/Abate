
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17550<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17550<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode17551 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17551<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17551 for Packet17551<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
