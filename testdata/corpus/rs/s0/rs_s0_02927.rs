
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode10780 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10780<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10780 for Packet10780<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome10781<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10781<T> {
    pub value: T,
    pub tag: u32,
}
