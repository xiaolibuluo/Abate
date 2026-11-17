
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17070<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17070<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode17071 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17071<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17071 for Packet17071<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
