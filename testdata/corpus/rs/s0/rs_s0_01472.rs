
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode5390 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5390<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5390 for Packet5390<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome5391<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5391<T> {
    pub value: T,
    pub tag: u32,
}
